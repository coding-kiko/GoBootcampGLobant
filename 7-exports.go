/* In Go, what defines if a definition (function, type, etc) is public or private is simply the
case of it's first letter. Supposing that it starts with a lower case letter, its definition
will be invisible outside of the package. In contrast, definitions starting will upper case
are perfectly usable when importing a package.

The clasic example is the pi implementation inside the math library: */

import "math"

pi := math.Pi // this works

notPi := math.pi // this doesn't