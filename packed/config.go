package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GAo1f8VwIAE+Bk4GZLz89Iy0/VLi3L0SvJzc0JDWBkYz0nLxuf0+fodNeBpuf5Vc82m1or+L35Be109ZMwXdTrXqZx6GdjMonPvXPL5mnnqUxqO7iovmeU8xXLzcvlUF8aEqjLfHt0yt+ebG8/+rGF3ORLIv+fhru9xt7QnBp9vj8pX+Va7R3J5p12SXTvP7FaVPQVaf5esj9C4eW3nuuvt9n9WtDe+vcT/c9HCW0fcVp9ZoDFt1Xx1trXh/g/zEmavvl9zWDpvg85/hcsJiy79+vPigdHaDXceWjMw/P8f4M3OYc7Q37uRgYFBnpGBAeZtBgYNA1Rvs8G9DfatjoxsPEgzspIAb0YmEWZEqCEbDAo1GFjSCCJxhSHCFOyOgAABhv+O3+GmIDmJlQ0kzcTAxNDMwMBgwgjiAQIAAP//1g4bTs0BAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
