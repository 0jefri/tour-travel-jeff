package repository

import (
	"database/sql"

	"github.com/golang-generic/model"
)

type PlanRepository interface {
	FindPlansByTourID(tourID int) ([]model.Plan, map[int][]model.Benefit, error)
}

type planRepository struct {
	db *sql.DB
}

func NewPlanRepository(db *sql.DB) PlanRepository {
	return &planRepository{db: db}
}

func (r *planRepository) FindPlansByTourID(tourID int) ([]model.Plan, map[int][]model.Benefit, error) {
	query := `
		SELECT p.id, p.tour_id, p.title, p.description, b.name AS benefit_name
		FROM plans p
		LEFT JOIN benefits b ON p.id = b.plan_id
		WHERE p.tour_id = $1
		ORDER BY p.id, b.id;
	`

	rows, err := r.db.Query(query, tourID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	plans := []model.Plan{}
	benefits := make(map[int][]model.Benefit)

	for rows.Next() {
		var plan model.Plan
		var benefitName sql.NullString

		err := rows.Scan(&plan.ID, &plan.TourID, &plan.Title, &plan.Description, &benefitName)
		if err != nil {
			return nil, nil, err
		}

		if len(plans) == 0 || plans[len(plans)-1].ID != plan.ID {
			plans = append(plans, plan)
		}

		if benefitName.Valid {
			benefits[plan.ID] = append(benefits[plan.ID], model.Benefit{Name: benefitName.String})
		}
	}

	for _, plan := range plans {
		if _, exists := benefits[plan.ID]; !exists {
			benefits[plan.ID] = []model.Benefit{}
		}
	}

	return plans, benefits, nil
}
