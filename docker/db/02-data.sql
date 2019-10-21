-- users 
INSERT INTO uuser (name) VALUES 
  ('Sadmin'), 
  ('Liliread'),
  ('Moderator')
;

-- all routes
INSERT INTO api_route (url, method) VALUES
  ('/people', 'GET'),
  ('/people', 'POST'),
  ('/people/:id', 'PUT'),
  ('/people/:id', 'DELETE')
;

-- all roles ( not users )
INSERT INTO api_role (name) VALUES
  ('reader'),
  ('updater'),
  ('deleter'),
  ('admin'),
  ('guest')
;

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
    FOREIGN KEY (user_id) REFERENCES uuser(user_id),
    CONSTRAINT api_role_user_api_role_id_user_id_key PRIMARY KEY (api_role_id, user_id)
);

-- link between route and role
INSERT INTO api_route_role (api_route_id, api_role_id) VALUES
  ((SELECT api_route_id FROM api_route WHERE url = '/people' and method = 'GET'), (SELECT api_role_id FROM api_role WHERE name = 'reader')),
  ((SELECT api_route_id FROM api_route WHERE url = '/people' and method = 'POST'), (SELECT api_role_id FROM api_role WHERE name = 'updater')),
  ((SELECT api_route_id FROM api_route WHERE url = '/people/:id' and method = 'PUT'), (SELECT api_role_id FROM api_role WHERE name = 'updater')),
  ((SELECT api_route_id FROM api_route WHERE url = '/people/:id' and method = 'DELETE'), (SELECT api_role_id FROM api_role WHERE name = 'deleter'))
;

-- link between role & role ( parent gives his rigths to child )
INSERT INTO api_role_role (parent_api_role_id, child_api_role_id) VALUES
  ((SELECT api_role_id FROM api_role WHERE name = 'reader'), (SELECT api_role_id FROM api_role WHERE name = 'admin')),
  ((SELECT api_role_id FROM api_role WHERE name = 'updater'), (SELECT api_role_id FROM api_role WHERE name = 'admin')),
  ((SELECT api_role_id FROM api_role WHERE name = 'deleter'), (SELECT api_role_id FROM api_role WHERE name = 'admin')),
  ((SELECT api_role_id FROM api_role WHERE name = 'reader'), (SELECT api_role_id FROM api_role WHERE name = 'guest'))
;

-- link between role & users
INSERT INTO api_role_user (api_role_id, user_id) VALUES
  ((SELECT api_role_id FROM api_role WHERE name = 'admin'),(SELECT user_id FROM uuser WHERE name = 'Sadmin')),
  ((SELECT api_role_id FROM api_role WHERE name = 'guest'),(SELECT user_id FROM uuser WHERE name = 'Liliread')),
  ((SELECT api_role_id FROM api_role WHERE name = 'deleter'),(SELECT user_id FROM uuser WHERE name = 'Moderator'))
;