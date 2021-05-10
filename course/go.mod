module "course"

go 1.15

require (
	github.com/rs/zerolog v1.21.0
	geometry v0.0.0
)
replace (
	geometry v0.0.0 => ../geometry
)