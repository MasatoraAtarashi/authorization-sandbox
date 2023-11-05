package app.rbac

default allow = false

allow {
    user_is_admin
}

user_is_admin {
    input.roles[_] = "admin"
}

role[p] {
    r := input.roles[_]
    p = data.role_permissions[r][_]
}

# GET /articles
allow {
    input.method = "GET"
    input.path = ["articles"]
    has_permission(input.roles, "article.list")
}

# GET /articles/:articleID
allow {
	some articleID
	input.method = "GET"
	input.path = ["articles", articleID]
	has_permission(input.roles, "article.get")
}

# POST /articles
allow {
	input.method = "POST"
	input.path = ["articles"]
	has_permission(input.roles, "article.create")
}

# PUT /articles/:articleID
allow {
	some articleID
	input.method = "PUT"
	input.path = ["articles", articleID]
	has_permission(input.roles, "article.update")
}

# DELETE /articles/:articleID
allow {
    input.method = "DELETE"
    input.path = ["articles", articleID]
    has_permission(input.roles, "article.delete")
}

has_permission(roles, p) {
    r := roles[_]
    data.role_permissions[r][_] = p
}