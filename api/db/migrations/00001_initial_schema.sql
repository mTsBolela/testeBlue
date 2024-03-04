CREATE TABLE IF NOT EXISTS public.categories (
id uuid NOT NULL,
description varchar NOT NULL,
type varchar NOT NULL,
status varchar(8) NOT NULL, -- active, inactive
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT categories_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.persons (
id uuid NOT NULL,
customer boolean NOT NULL,
provider boolean NOT NULL,
document varchar(18) NOT NULL,
name varchar NOT NULL,
cep varchar(10) NOT NULL,
address varchar NOT NULL,
state varchar NOT NULL,
city varchar NOT NULL,
complement varchar NOT NULL,
status varchar(8) NOT NULL, -- active, inactive
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
CONSTRAINT persons_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.financials (
id uuid NOT NULL,
type varchar(7) NOT NULL, -- payment / receipt
category_id uuid NOT NULL,
person_id uuid NOT NULL,
document varchar NOT NULL,
discharge_value numeric(15,2) NOT NULL,
document_value numeric(15,2) NOT NULL,
description text NOT NULL,
due_date DATE NOT NULL,
discharge_date DATE,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP,
status VARCHAR(8), --pending, paidout, canceled
CONSTRAINT financials_pkey PRIMARY KEY (id)
);
	 
ALTER TABLE public.financials
ADD CONSTRAINT fk_financials_categories
FOREIGN KEY (category_id)
REFERENCES categories(id);

ALTER TABLE public.financials
ADD CONSTRAINT fk_financials_persons
FOREIGN KEY (person_id)
REFERENCES persons(id);

CREATE INDEX  idx_description ON categories (description);
CREATE INDEX  idx_name ON persons (name);
CREATE INDEX  idx_financials_due ON financials (due_date, category_id, person_id);
CREATE INDEX  idx_financials_dischange ON financials (dischange_date, category_id, person_id);