--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.groups (
    id integer NOT NULL,
    name text,
    uuid text
);


ALTER TABLE public.groups OWNER TO postgres;

--
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.groups_id_seq OWNER TO postgres;

--
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- Name: groupstotasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.groupstotasks (
    id integer NOT NULL,
    group_id text,
    task_id text
);


ALTER TABLE public.groupstotasks OWNER TO postgres;

--
-- Name: groupstotasks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.groupstotasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.groupstotasks_id_seq OWNER TO postgres;

--
-- Name: groupstotasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.groupstotasks_id_seq OWNED BY public.groupstotasks.id;


--
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    id integer NOT NULL,
    name text,
    uuid text
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tasks_id_seq OWNER TO postgres;

--
-- Name: tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.tasks.id;


--
-- Name: taskstotimeframes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.taskstotimeframes (
    id integer NOT NULL,
    task_id text,
    timeframe_id text
);


ALTER TABLE public.taskstotimeframes OWNER TO postgres;

--
-- Name: taskstotimeframes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.taskstotimeframes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.taskstotimeframes_id_seq OWNER TO postgres;

--
-- Name: taskstotimeframes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.taskstotimeframes_id_seq OWNED BY public.taskstotimeframes.id;


--
-- Name: timeframes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.timeframes (
    id integer NOT NULL,
    start text,
    stop text,
    uuid text
);


ALTER TABLE public.timeframes OWNER TO postgres;

--
-- Name: timeframes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.timeframes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.timeframes_id_seq OWNER TO postgres;

--
-- Name: timeframes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.timeframes_id_seq OWNED BY public.timeframes.id;


--
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- Name: groupstotasks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groupstotasks ALTER COLUMN id SET DEFAULT nextval('public.groupstotasks_id_seq'::regclass);


--
-- Name: tasks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);


--
-- Name: taskstotimeframes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.taskstotimeframes ALTER COLUMN id SET DEFAULT nextval('public.taskstotimeframes_id_seq'::regclass);


--
-- Name: timeframes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.timeframes ALTER COLUMN id SET DEFAULT nextval('public.timeframes_id_seq'::regclass);


--
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.groups (id, name, uuid) FROM stdin;
\.


--
-- Data for Name: groupstotasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.groupstotasks (id, group_id, task_id) FROM stdin;
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tasks (id, name, uuid) FROM stdin;
\.


--
-- Data for Name: taskstotimeframes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.taskstotimeframes (id, task_id, timeframe_id) FROM stdin;
\.


--
-- Data for Name: timeframes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.timeframes (id, start, stop, uuid) FROM stdin;
\.


--
-- Name: groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.groups_id_seq', 1, false);


--
-- Name: groupstotasks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.groupstotasks_id_seq', 1, false);


--
-- Name: tasks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tasks_id_seq', 1, false);


--
-- Name: taskstotimeframes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.taskstotimeframes_id_seq', 1, false);


--
-- Name: timeframes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.timeframes_id_seq', 1, false);


--
-- Name: groups groups_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_name_key UNIQUE (name);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- Name: groups groups_uuid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_uuid_key UNIQUE (uuid);


--
-- Name: groupstotasks groupstotasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groupstotasks
    ADD CONSTRAINT groupstotasks_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_name_key UNIQUE (name);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_uuid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_uuid_key UNIQUE (uuid);


--
-- Name: taskstotimeframes taskstotimeframes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.taskstotimeframes
    ADD CONSTRAINT taskstotimeframes_pkey PRIMARY KEY (id);


--
-- Name: timeframes timeframes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.timeframes
    ADD CONSTRAINT timeframes_pkey PRIMARY KEY (id);


--
-- Name: timeframes timeframes_uuid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.timeframes
    ADD CONSTRAINT timeframes_uuid_key UNIQUE (uuid);


--
-- PostgreSQL database dump complete
--

