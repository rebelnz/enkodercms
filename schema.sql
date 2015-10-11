	CREATE TABLE IF NOT EXISTS account (
    id serial PRIMARY KEY,
    email character varying(255) NOT NULL DEFAULT '',
    password character varying(255) NOT NULL DEFAULT '',
    firstname character varying(255) NOT NULL DEFAULT '',
    lastname character varying(255) NOT NULL DEFAULT '',
    avatar character varying(255) NOT NULL DEFAULT '',
    priv bigint DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS subscriber (
    id serial PRIMARY KEY,
    randomid character varying(255) NOT NULL DEFAULT '',
    email character varying(255) NOT NULL DEFAULT '',
    firstname character varying(255) NOT NULL DEFAULT '',
    lastname character varying(255) NOT NULL DEFAULT '',
    newsletter bigint DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS settings (
    id serial PRIMARY KEY,
    email character varying(255) NOT NULL DEFAULT '',
    address character varying(255) NOT NULL DEFAULT '',
    street character varying(255) NOT NULL DEFAULT '',
    suburb character varying(255) NOT NULL DEFAULT '',
    city character varying(255) NOT NULL DEFAULT '',
    code character varying(255) NOT NULL DEFAULT '',
    contact character varying(255) NOT NULL DEFAULT '',
    twitter character varying(255) NOT NULL DEFAULT '',
    facebook character varying(255) NOT NULL DEFAULT '',
    linkedin character varying(255) NOT NULL DEFAULT '',
    description character varying(255) NOT NULL DEFAULT '',
    keywords character varying(255) NOT NULL DEFAULT '',
    ganalytics character varying(255) NOT NULL DEFAULT '',
    smtp character varying(255) NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS page (
    id serial PRIMARY KEY,
    slug character varying(255) NOT NULL DEFAULT '',
    title character varying(255) NOT NULL DEFAULT '',
    content text NOT NULL DEFAULT '',
    metatags text NOT NULL DEFAULT '',
    "order" bigint DEFAULT 0,
    priv bigint DEFAULT 0,
    status bigint DEFAULT 0,
    parent bigint DEFAULT 0,
    navtype character varying(255) NOT NULL DEFAULT 'topnav',
    navorder bigint DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS widget (
    id serial PRIMARY KEY,
    title character varying(255) NOT NULL DEFAULT '',
    content text NOT NULL DEFAULT '',
    url character varying(255) NOT NULL DEFAULT '',
    filename character varying(255) NOT NULL DEFAULT '',
    widgetsize character varying(255) NOT NULL DEFAULT 'col-1',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS media (
    id serial PRIMARY KEY,
    title character varying(255) NOT NULL DEFAULT '',
    filename character varying(255) NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS coords (
    latitude float NOT NULL,
    longitude float NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());

    CREATE TABLE IF NOT EXISTS page_media (
    page_id integer REFERENCES page (id) ON UPDATE CASCADE ON DELETE CASCADE,
    media_id integer REFERENCES media (id) ON UPDATE CASCADE,
    CONSTRAINT page_media_pkey PRIMARY KEY (page_id, media_id));

    CREATE TABLE IF NOT EXISTS page_widget (
    page_id    integer REFERENCES page (id) ON UPDATE CASCADE ON DELETE CASCADE,
    widget_id integer REFERENCES widget (id) ON UPDATE CASCADE,
    CONSTRAINT page_widget_pkey PRIMARY KEY (page_id, widget_id));

    CREATE TABLE IF NOT EXISTS slider (
    id serial PRIMARY KEY,
    title character varying(255) NOT NULL DEFAULT '',
    content text NOT NULL DEFAULT '',
    filename character varying(255) NOT NULL DEFAULT '',
    url character varying(255) NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());
    
    CREATE TABLE IF NOT EXISTS post (
    id serial PRIMARY KEY,
    account_id integer REFERENCES account (id),
    slug text NOT NULL DEFAULT '',
    title character varying(255) NOT NULL DEFAULT '',
    content text NOT NULL DEFAULT '',
    filename character varying(255) NOT NULL DEFAULT '',
    tags text NOT NULL DEFAULT '',
    priv bigint DEFAULT 0,
    status bigint DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now());


    -- CREATE TABLE IF NOT EXISTS item (
    -- id serial PRIMARY KEY,
    -- account_id integer REFERENCES account (id),
    -- slug character varying(255) NOT NULL DEFAULT '',
    -- parent bigint DEFAULT 0,
    -- navtype character varying(255) NOT NULL DEFAULT 'topnav',
    -- navorder bigint DEFAULT 0,
    -- "type" character varying(255) NOT NULL DEFAULT 'page',
    -- title character varying(255) NOT NULL DEFAULT '',
    -- content text NOT NULL DEFAULT '',
    -- link character varying(500) NOT NULL DEFAULT '',
    -- filename character varying(255) NOT NULL DEFAULT '',
    -- widgetsize character varying(255) NOT NULL DEFAULT 'col-1',
    -- tags text NOT NULL DEFAULT '',
    -- metatags text NOT NULL DEFAULT '',
    -- "order" bigint DEFAULT 0,
    -- priv bigint DEFAULT 0,
    -- status bigint DEFAULT 0,
    -- created_at timestamp NOT NULL DEFAULT now(),
    -- updated_at timestamp NOT NULL DEFAULT now());
