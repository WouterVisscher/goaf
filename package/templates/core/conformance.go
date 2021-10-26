package coretemplates

import(
  "oaf-server/package/models"
)

// Function definition for conformance classes
type RenderConformanceFunc func (*models.Webcontext, *models.Conformanceclasses)

// Interface definition for conformance classes
type RenderConformanceType interface {
  RenderConformance(context *models.Webcontext, conformanceClasses *models.Conformanceclasses)
}

// Transforms a renderconformance function into a renderconformance object
func NewRenderConformanceType(fun RenderConformanceFunc) RenderConformanceType {
  return &renderConformanceType{
    renderConformanceFunc: fun,
  }
}

// Internal

type renderConformanceType struct {
  renderConformanceFunc RenderConformanceFunc
}

func (object *renderConformanceType) RenderConformance(context *models.Webcontext, conformanceClasses *models.Conformanceclasses) {
  object.renderConformanceFunc(context, conformanceClasses)
}
