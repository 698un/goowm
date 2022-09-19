module main

go 1.18

//CONFIG unit
replace cnfg => /cnfg
//require cnfg v0.0.0-00010101000000-000000000000 // indirect


//modules of math
replace mymath => ../mymath
require mymath v0.0.0-00010101000000-000000000000 // indirect

//modules of NET
replace netcontrol => /netcontrol
require netcontrol v0.0.0-00010101000000-000000000000 // indirect

