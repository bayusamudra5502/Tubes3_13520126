--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2 (Ubuntu 14.2-1ubuntu1)

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

DROP DATABASE IF EXISTS dnachecker;
--
-- Name: dnachecker; Type: DATABASE; Schema: -; Owner: root
--

CREATE DATABASE dnachecker WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE dnachecker OWNER TO root;

\connect dnachecker

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
-- Name: pemeriksaan; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.pemeriksaan (
    id_pemeriksaan integer NOT NULL,
    timestamp_created timestamp without time zone DEFAULT now(),
    nama_pasien character varying(100) NOT NULL,
    dna_sampel text NOT NULL,
    id_penyakit integer NOT NULL,
    similarity double precision NOT NULL
);


ALTER TABLE public.pemeriksaan OWNER TO root;

--
-- Name: pemeriksaan_id_pemeriksaan_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.pemeriksaan_id_pemeriksaan_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pemeriksaan_id_pemeriksaan_seq OWNER TO root;

--
-- Name: pemeriksaan_id_pemeriksaan_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.pemeriksaan_id_pemeriksaan_seq OWNED BY public.pemeriksaan.id_pemeriksaan;


--
-- Name: pemeriksaan_id_penyakit_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.pemeriksaan_id_penyakit_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pemeriksaan_id_penyakit_seq OWNER TO root;

--
-- Name: pemeriksaan_id_penyakit_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.pemeriksaan_id_penyakit_seq OWNED BY public.pemeriksaan.id_penyakit;


--
-- Name: penyakit; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.penyakit (
    id_penyakit integer NOT NULL,
    nama character varying(100) NOT NULL,
    timestamp_created timestamp without time zone DEFAULT now(),
    dna text NOT NULL
);


ALTER TABLE public.penyakit OWNER TO root;

--
-- Name: penyakit_id_penyakit_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.penyakit_id_penyakit_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.penyakit_id_penyakit_seq OWNER TO root;

--
-- Name: penyakit_id_penyakit_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.penyakit_id_penyakit_seq OWNED BY public.penyakit.id_penyakit;


--
-- Name: pemeriksaan id_pemeriksaan; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.pemeriksaan ALTER COLUMN id_pemeriksaan SET DEFAULT nextval('public.pemeriksaan_id_pemeriksaan_seq'::regclass);


--
-- Name: pemeriksaan id_penyakit; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.pemeriksaan ALTER COLUMN id_penyakit SET DEFAULT nextval('public.pemeriksaan_id_penyakit_seq'::regclass);


--
-- Name: penyakit id_penyakit; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.penyakit ALTER COLUMN id_penyakit SET DEFAULT nextval('public.penyakit_id_penyakit_seq'::regclass);


--
-- Data for Name: pemeriksaan; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.pemeriksaan (id_pemeriksaan, timestamp_created, nama_pasien, dna_sampel, id_penyakit, similarity) FROM stdin;
\.


--
-- Data for Name: penyakit; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.penyakit (id_penyakit, nama, timestamp_created, dna) FROM stdin;
\.


--
-- Name: pemeriksaan_id_pemeriksaan_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.pemeriksaan_id_pemeriksaan_seq', 1, false);


--
-- Name: pemeriksaan_id_penyakit_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.pemeriksaan_id_penyakit_seq', 1, false);


--
-- Name: penyakit_id_penyakit_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.penyakit_id_penyakit_seq', 1, false);


--
-- Name: pemeriksaan pemeriksaan_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.pemeriksaan
    ADD CONSTRAINT pemeriksaan_pk PRIMARY KEY (id_pemeriksaan);


--
-- Name: penyakit penyakit_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.penyakit
    ADD CONSTRAINT penyakit_pk PRIMARY KEY (id_penyakit);


--
-- Name: pemeriksaan pemeriksaan_penyakit_id_penyakit_fk; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.pemeriksaan
    ADD CONSTRAINT pemeriksaan_penyakit_id_penyakit_fk FOREIGN KEY (id_penyakit) REFERENCES public.penyakit(id_penyakit) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

