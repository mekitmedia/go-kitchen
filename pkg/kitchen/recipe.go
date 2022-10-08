package kitchen

import "context"

type RecipeStep interface {
	Run(ctx context.Context, ingredients interface{}) error
}

type Recipe struct {
	step []RecipeStep
}

func (r *Recipe) AddStep(stage RecipeStep) {
	r.step = append(r.step, stage)
}

func (r *Recipe) Make(ctx context.Context, ingredients interface{}) error {
	for _, step := range r.step {
		err := step.Run(ctx, ingredients)
		if err != nil {
			return err
		}
	}

	return nil
}
