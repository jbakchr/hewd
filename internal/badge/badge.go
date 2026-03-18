package badge

import (
	"fmt"
)

// Color returns a color name based on score according to typical badge conventions.
func Color(score int) string {
	switch {
	case score >= 90:
		return "#4c1" // bright green
	case score >= 75:
		return "#97CA00" // green
	case score >= 60:
		return "#dfb317" // yellow
	case score >= 40:
		return "#fe7d37" // orange
	default:
		return "#e05d44" // red
	}
}

// Generate produces an SVG badge for a given score.
func Generate(score int) string {
	color := Color(score)
	return fmt.Sprintf(`
<svg xmlns="http://www.w3.org/2000/svg" width="150" height="20" role="img" aria-label="hewd score: %d">
  <linearGradient id="smooth" x2="0" y2="100%%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>
  <rect rx="3" width="150" height="20" fill="#555"/>
  <rect rx="3" x="85" width="65" height="20" fill="%s"/>
  <rect rx="3" width="150" height="20" fill="url(#smooth)"/>
  <g fill="#fff" text-anchor="middle"
     font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text x="42.5" y="15">hewd score</text>
    <text x="117.5" y="15">%d</text>
  </g>
</svg>
`, score, color, score)
}
