/**
 * USER 
 */
CREATE TABLE users
(
    user_id serial PRIMARY KEY,
    name text NOT NULL,
    CONSTRAINT user_name UNIQUE (name)
);

/**
 * RBAC Implementation
 */
CREATE TYPE enum_http_method AS ENUM ('GET', 'POST', 'PUT', 'DELETE');

CREATE TABLE IF NOT EXISTS "api_route" (
    "api_route_id" serial PRIMARY KEY,
    "url" text NOT NULL,
    "method" enum_http_method NOT NULL,
    CONSTRAINT api_route_url_method_key UNIQUE (url, method)
);

CREATE TABLE IF NOT EXISTS "api_role" (
    "api_role_id" serial PRIMARY KEY,
    "name" text NOT NULL,
    CONSTRAINT api_role_name_key UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS "api_role_role" (
    "parent_api_role_id" integer NOT NULL,
    "child_api_role_id" integer NOT NULL,
    FOREIGN KEY (parent_api_role_id) REFERENCES api_role(api_role_id),
    FOREIGN KEY (child_api_role_id) REFERENCES api_role(api_role_id),
    CONSTRAINT api_role_role_parent_api_role_id_child_api_role_id_key PRIMARY KEY (parent_api_role_id, child_api_role_id)
);

CREATE TABLE IF NOT EXISTS "api_route_role" (
    api_route_id integer NOT NULL,
    api_role_id integer NOT NULL,
    FOREIGN KEY (api_route_id) REFERENCES api_route(api_route_id),
    FOREIGN KEY (api_role_id) REFERENCES api_role(api_role_id),
    CONSTRAINT api_route_role_api_route_id_api_role_id_key PRIMARY KEY (api_route_id, api_role_id)
);

CREATE TABLE IF NOT EXISTS "api_role_user" (
    api_role_id integer NOT NULL,
    user_id integer NOT NULL,
    FOREIGN KEY (api_role_id) REFERENCES api_role(api_role_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT api_role_user_api_role_id_user_id_key PRIMARY KEY (api_role_id, user_id)
);
