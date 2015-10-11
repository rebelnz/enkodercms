--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: account; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE account (
    id integer NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    password character varying(255) DEFAULT ''::character varying NOT NULL,
    firstname character varying(255) DEFAULT ''::character varying NOT NULL,
    lastname character varying(255) DEFAULT ''::character varying NOT NULL,
    avatar character varying(255) DEFAULT ''::character varying NOT NULL,
    priv bigint DEFAULT 0,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.account OWNER TO enkoder;

--
-- Name: account_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE account_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.account_id_seq OWNER TO enkoder;

--
-- Name: account_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE account_id_seq OWNED BY account.id;


--
-- Name: coords; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE coords (
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.coords OWNER TO enkoder;

--
-- Name: media; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE media (
    id integer NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    filename character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.media OWNER TO enkoder;

--
-- Name: media_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE media_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.media_id_seq OWNER TO enkoder;

--
-- Name: media_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE media_id_seq OWNED BY media.id;


--
-- Name: page; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE page (
    id integer NOT NULL,
    slug character varying(255) DEFAULT ''::character varying NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    content text DEFAULT ''::text NOT NULL,
    metatags text DEFAULT ''::text NOT NULL,
    "order" bigint DEFAULT 0,
    priv bigint DEFAULT 0,
    status bigint DEFAULT 0,
    parent bigint DEFAULT 0,
    navtype character varying(255) DEFAULT 'topnav'::character varying NOT NULL,
    navorder bigint DEFAULT 0,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.page OWNER TO enkoder;

--
-- Name: page_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE page_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.page_id_seq OWNER TO enkoder;

--
-- Name: page_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE page_id_seq OWNED BY page.id;


--
-- Name: page_media; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE page_media (
    page_id integer NOT NULL,
    media_id integer NOT NULL
);


ALTER TABLE public.page_media OWNER TO enkoder;

--
-- Name: page_widget; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE page_widget (
    page_id integer NOT NULL,
    widget_id integer NOT NULL
);


ALTER TABLE public.page_widget OWNER TO enkoder;

--
-- Name: post; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE post (
    id integer NOT NULL,
    account_id integer,
    slug text DEFAULT ''::text NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    content text DEFAULT ''::text NOT NULL,
    filename character varying(255) DEFAULT ''::character varying NOT NULL,
    tags text DEFAULT ''::text NOT NULL,
    priv bigint DEFAULT 0,
    status bigint DEFAULT 0,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.post OWNER TO enkoder;

--
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO enkoder;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE post_id_seq OWNED BY post.id;


--
-- Name: settings; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE settings (
    id integer NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    address character varying(255) DEFAULT ''::character varying NOT NULL,
    street character varying(255) DEFAULT ''::character varying NOT NULL,
    suburb character varying(255) DEFAULT ''::character varying NOT NULL,
    city character varying(255) DEFAULT ''::character varying NOT NULL,
    code character varying(255) DEFAULT ''::character varying NOT NULL,
    contact character varying(255) DEFAULT ''::character varying NOT NULL,
    twitter character varying(255) DEFAULT ''::character varying NOT NULL,
    facebook character varying(255) DEFAULT ''::character varying NOT NULL,
    linkedin character varying(255) DEFAULT ''::character varying NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    keywords character varying(255) DEFAULT ''::character varying NOT NULL,
    ganalytics character varying(255) DEFAULT ''::character varying NOT NULL,
    smtp character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.settings OWNER TO enkoder;

--
-- Name: settings_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE settings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.settings_id_seq OWNER TO enkoder;

--
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE settings_id_seq OWNED BY settings.id;


--
-- Name: slider; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE slider (
    id integer NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    content text DEFAULT ''::text NOT NULL,
    filename character varying(255) DEFAULT ''::character varying NOT NULL,
    url character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.slider OWNER TO enkoder;

--
-- Name: slider_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE slider_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.slider_id_seq OWNER TO enkoder;

--
-- Name: slider_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE slider_id_seq OWNED BY slider.id;


--
-- Name: subscriber; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE subscriber (
    id integer NOT NULL,
    randomid character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    firstname character varying(255) DEFAULT ''::character varying NOT NULL,
    lastname character varying(255) DEFAULT ''::character varying NOT NULL,
    newsletter bigint DEFAULT 0,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.subscriber OWNER TO enkoder;

--
-- Name: subscriber_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE subscriber_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscriber_id_seq OWNER TO enkoder;

--
-- Name: subscriber_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE subscriber_id_seq OWNED BY subscriber.id;


--
-- Name: widget; Type: TABLE; Schema: public; Owner: enkoder; Tablespace: 
--

CREATE TABLE widget (
    id integer NOT NULL,
    title character varying(255) DEFAULT ''::character varying NOT NULL,
    content text DEFAULT ''::text NOT NULL,
    url character varying(255) DEFAULT ''::character varying NOT NULL,
    filename character varying(255) DEFAULT ''::character varying NOT NULL,
    widgetsize character varying(255) DEFAULT 'col-1'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.widget OWNER TO enkoder;

--
-- Name: widget_id_seq; Type: SEQUENCE; Schema: public; Owner: enkoder
--

CREATE SEQUENCE widget_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.widget_id_seq OWNER TO enkoder;

--
-- Name: widget_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: enkoder
--

ALTER SEQUENCE widget_id_seq OWNED BY widget.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY account ALTER COLUMN id SET DEFAULT nextval('account_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY media ALTER COLUMN id SET DEFAULT nextval('media_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY page ALTER COLUMN id SET DEFAULT nextval('page_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY post ALTER COLUMN id SET DEFAULT nextval('post_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY settings ALTER COLUMN id SET DEFAULT nextval('settings_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY slider ALTER COLUMN id SET DEFAULT nextval('slider_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY subscriber ALTER COLUMN id SET DEFAULT nextval('subscriber_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY widget ALTER COLUMN id SET DEFAULT nextval('widget_id_seq'::regclass);


--
-- Data for Name: account; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY account (id, email, password, firstname, lastname, avatar, priv, created_at, updated_at) FROM stdin;
1	dev@enkoder.com.au	$2a$10$WFtTOs8pbEK4mYuTvnNPj.3212tn.Xqr29HmEXbCecRJ5mRF0RSBO	Dev	Enkoder	1427873084avatar.jpg	3	2015-03-29 11:43:01.485996	2015-04-01 18:24:44.455651
\.


--
-- Name: account_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('account_id_seq', 1, true);


--
-- Data for Name: coords; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY coords (latitude, longitude, created_at, updated_at) FROM stdin;
-37.8145305399571754	144.963226318359375	2015-03-29 11:43:01.513581	2015-04-01 19:30:44.752346
\.


--
-- Data for Name: media; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY media (id, title, filename, created_at, updated_at) FROM stdin;
\.


--
-- Name: media_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('media_id_seq', 1, false);


--
-- Data for Name: page; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY page (id, slug, title, content, metatags, "order", priv, status, parent, navtype, navorder, created_at, updated_at) FROM stdin;
1	home	Home			0	0	0	0	topnav	0	2015-03-29 11:43:01.495606	2015-03-29 11:43:01.495606
3	contact	Contact			0	0	1	0	topnav	2	2015-03-29 11:43:01.502272	2015-03-31 15:48:21.857118
6	enkoder-cms	Enkoder CMS	The Enkoder CMS provides a simple and intuitive way for clients to update and add new content and media.The base system includes analytics and user management by default. The inbuilt "News" and subscribers functionality is critical for digital marketing and the SEO settings are all accessible from the administration console.<br><br>Full customisation is possible and is only limited by your imagination. The Admin area is also built to be responsive and contains a full manual on how to configure your website.<br>	CMS, responsive web design, web development, custom digital solutions	0	0	1	0	topnav	3	2015-04-01 18:20:21.176083	2015-04-01 18:43:16.207351
2	map	Map			0	0	1	0	footernav	1	2015-03-29 11:43:01.49924	2015-03-31 21:16:08.7813
4	about	About	<h4>Open Source</h4>Enkoder uses open source software such as Linux, Postgres, Go and Debian in order to offer competitive pricing. Our clients directly benefit from updates, enhancements and new functionality as implemented. This is especially important in regards to security patches and updates.\r\n\r\n<br><br><h4>Responsive Design\r\n</h4>Design and development is targeted to mobile first to leverage the ever increasing mobile market. Our development stack ensures we can stay current with all the developments in regards to the mobile landscape. This also means our clients can leverage mobile technologies such as Geo-location.\r\n<br>\r\n<br><h4>Custom CMS\r\n</h4>The Enkoder Content Management System has been developed to be intuitive and user friendly. Any customisations or extensions can be implemented at any stage. The base system includes and analytics dashboard, user management, configuration settings, map functionality and a basic contact form.\r\n\r\n<br><br><h4>Support 24/7\r\n</h4>We pride ourselves on excellent customer relationships and are able to provide support at all stages of development. We are happy provide technical support.		0	0	1	0	topnav	1	2015-03-31 15:48:36.039195	2015-04-01 19:19:05.470586
\.


--
-- Name: page_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('page_id_seq', 6, true);


--
-- Data for Name: page_media; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY page_media (page_id, media_id) FROM stdin;
\.


--
-- Data for Name: page_widget; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY page_widget (page_id, widget_id) FROM stdin;
6	3
6	2
4	4
4	5
\.


--
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY post (id, account_id, slug, title, content, filename, tags, priv, status, created_at, updated_at) FROM stdin;
1	1	news/2015-Apr-01/balgowlah-physio	Balgowlah Physio	<h3>Balgowlah Physio Website</h3>Sydney based Physiotherapists - Balgowlah Physio - required a revamp of their existing website. The new site is built to be responsive and looks great across all devices. The owner is able to sign in and create or update information as needed as well as provide new and interesting information and advice to potential customers.<br><br>The job included logo design and a full re-branding of their website. The home page boldly features a map and their contact information.<br><br><a title="Balgowlah Physio Site" href="http://balgowlahphysio.com">Balgowlah Physio Site</a><br>	1427875411balgowlah-physio.jpg	Portfolio, responsive, websites	0	1	2015-04-01 10:58:54.820948	2015-04-01 19:10:48.377842
\.


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('post_id_seq', 1, true);


--
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY settings (id, email, address, street, suburb, city, code, contact, twitter, facebook, linkedin, description, keywords, ganalytics, smtp, created_at, updated_at) FROM stdin;
1	dev@enkoder.com.au				Melbourne 	3000	0449726735				Melbourne Web Design and Development	web design melbourne, australia, digital marketing, web based applications	386498933990-3m58l6kfo8km6geh49ihsargp3msamus.apps.googleusercontent.com		2015-03-29 11:43:01.507742	2015-04-01 18:45:00.443926
\.


--
-- Name: settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('settings_id_seq', 1, true);


--
-- Data for Name: slider; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY slider (id, title, content, filename, url, created_at, updated_at) FROM stdin;
\.


--
-- Name: slider_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('slider_id_seq', 1, false);


--
-- Data for Name: subscriber; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY subscriber (id, randomid, email, firstname, lastname, newsletter, created_at, updated_at) FROM stdin;
2	6CVeHlKmRlMhSf39jfGkljSgJ3u9dSyZ	rebel@gmail.com			1	2015-04-01 13:45:49.979539	2015-04-01 13:45:49.979539
\.


--
-- Name: subscriber_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('subscriber_id_seq', 3, true);


--
-- Data for Name: widget; Type: TABLE DATA; Schema: public; Owner: enkoder
--

COPY widget (id, title, content, url, filename, widgetsize, created_at, updated_at) FROM stdin;
3	CMS Analytics	<h3>CMS Analytics</h3>From the administration section you are able to get a high level view of your Google analytics.<br><br>		1427873882enkoder-cms-admin-analytics.png	small	2015-04-01 18:38:03.003605	2015-04-01 18:39:18.737012
2	CMS Pages	<h3>CMS Content<br></h3><p>Adding or editing content to pages is simple and intuitive.<br></p>		1427873776enkoder-cms-admin-page.png	small	2015-04-01 18:36:16.890308	2015-04-01 18:39:47.277895
4	Enkoder CMS	<h3>Enkoder CMS\r\n</h3>Our built-in CMS is intuitive and simple to use. Settings and user management come as default and more customizations are possible.<br><br>		1427876148admin-pages.png	small	2015-04-01 19:15:48.248927	2015-04-01 19:16:15.850105
5	Devices	<h3>Mobile and Desktop</h3><p>All our design are designed to be responsive and look great across a range of devices. Whether your users are browsing by phone, tablet or desktop they will be able to access all the information.<br></p>		1427876335enkoder-devices.jpg	small	2015-04-01 19:18:55.089687	2015-04-01 19:18:55.089687
\.


--
-- Name: widget_id_seq; Type: SEQUENCE SET; Schema: public; Owner: enkoder
--

SELECT pg_catalog.setval('widget_id_seq', 5, true);


--
-- Name: account_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY account
    ADD CONSTRAINT account_pkey PRIMARY KEY (id);


--
-- Name: media_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY media
    ADD CONSTRAINT media_pkey PRIMARY KEY (id);


--
-- Name: page_media_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY page_media
    ADD CONSTRAINT page_media_pkey PRIMARY KEY (page_id, media_id);


--
-- Name: page_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY page
    ADD CONSTRAINT page_pkey PRIMARY KEY (id);


--
-- Name: page_widget_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY page_widget
    ADD CONSTRAINT page_widget_pkey PRIMARY KEY (page_id, widget_id);


--
-- Name: post_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);


--
-- Name: settings_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


--
-- Name: slider_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY slider
    ADD CONSTRAINT slider_pkey PRIMARY KEY (id);


--
-- Name: subscriber_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY subscriber
    ADD CONSTRAINT subscriber_pkey PRIMARY KEY (id);


--
-- Name: widget_pkey; Type: CONSTRAINT; Schema: public; Owner: enkoder; Tablespace: 
--

ALTER TABLE ONLY widget
    ADD CONSTRAINT widget_pkey PRIMARY KEY (id);


--
-- Name: page_media_media_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY page_media
    ADD CONSTRAINT page_media_media_id_fkey FOREIGN KEY (media_id) REFERENCES media(id) ON UPDATE CASCADE;


--
-- Name: page_media_page_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY page_media
    ADD CONSTRAINT page_media_page_id_fkey FOREIGN KEY (page_id) REFERENCES page(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: page_widget_page_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY page_widget
    ADD CONSTRAINT page_widget_page_id_fkey FOREIGN KEY (page_id) REFERENCES page(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: page_widget_widget_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY page_widget
    ADD CONSTRAINT page_widget_widget_id_fkey FOREIGN KEY (widget_id) REFERENCES widget(id) ON UPDATE CASCADE;


--
-- Name: post_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: enkoder
--

ALTER TABLE ONLY post
    ADD CONSTRAINT post_account_id_fkey FOREIGN KEY (account_id) REFERENCES account(id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

