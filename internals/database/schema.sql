create table if not exists company (
    id uuid primary key default gen_random_uuid(),
    ruc text not null,
    name text not null,
    employees smallint default 1,
    is_active boolean default true,
    created_at timestamp with time zone default now(),

    unique (ruc),
    unique (name)
);

