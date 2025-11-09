-- Create SysAdmin Account
INSERT INTO "accounts" (
    "id",
    "name",
    "slug",
    "logo",
    "domain",
    "domain_verified",
    "account_id",
    "created_at",
    "created_by",
    "updated_at",
    "updated_by",
    "deleted_at",
    "deleted_by"
)
VALUES (
    '00000000-0000-0000-0000-000000000000',
    'SwiftGeek',
    'accounts',
    'http://localhost:3000/logo.png',
    'localhost:8080',
    true,
    '00000000-0000-0000-0000-000000000000',
    NOW(),
    '00000000-0000-0000-0000-000000000000',
    NULL,
    NULL,
    NULL,
    NULL
);


-- Create Default Application
INSERT INTO "apps" (
    "id",
    "name",
    "description",
    "landing_url",
    "logo",
    "branding",
    "client_id",
    "client_secret",
    "redirect_uris",
    "jwt_algo",
    "jwt_secret",
    "jwt_lifetime",
    "refresh_token_lifetime",
    "permanent_callback",
    "permanent_error_callback",
    "account_id",
    "created_at",
    "created_by",
    "updated_at",
    "updated_by",
    "deleted_at",
    "deleted_by"
)
VALUES (
    '00000000-0000-0000-0000-000000000000',
    'AuthInfinity',
    'Quickly setup OAuth + OpenID Connect',
    'http://localhost:3000/',
    NULL,
    NULL,
    'authinfinity',
    'authinfinity-secret',
    '{"http://localhost:3000/authorize"}',
    'HS256',
    'myjwtsecret',
    '1d',
    '15d',
    'http://localhost:3000/authorize',
    'http://localhost:3000/error',
    '00000000-0000-0000-0000-000000000000',
    NOW(),
    '00000000-0000-0000-0000-000000000000',
    NULL,
    NULL,
    NULL,
    NULL
);
-- Create Sys Admin
INSERT INTO "users" (
    "id",
    "name",
    "email",
    "dp",
    "email_verified",
    "account_id",
    "created_at",
    "created_by",
    "updated_at",
    "updated_by",
    "deleted_at",
    "deleted_by"
)
VALUES (
    '00000000-0000-0000-0000-000000000000',
    'Sys Admin',
    'admin@swiftgeek.com',
    NULL,
    true,
    '00000000-0000-0000-0000-000000000000',
    NOW(),
    '00000000-0000-0000-0000-000000000000',
    NULL,
    NULL,
    NULL,
    NULL
);

-- Create Sys Admin Password 
INSERT INTO "passwords" (
    "id",
    "password",
    "user_id",
    "account_id",
    "created_at",
    "created_by",
    "updated_at",
    "updated_by",
    "deleted_at",
    "deleted_by"
)
VALUES (
    '00000000-0000-0000-0000-000000000000',
    '$2a$10$O9ZOlPzoo1aMrd3XkXYjD.zwCzmGvKVAdzCk9IedlmgkvVZkxLjcm', -- Test@1234
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000',
    NOW(),
    '00000000-0000-0000-0000-000000000000',
    NULL,
    NULL,
    NULL,
    NULL
);




-- TRUNCATE TABLE accounts, apps, email_verification_requests, passwords, reset_password_requests, sessions, users;

