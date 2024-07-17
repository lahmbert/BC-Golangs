create table entity_type (
	enty_id SERIAL,
	enty_name VARCHAR (25),
	constraint enty_id_pk primary key(enty_id),
	constraint entyy_name_uq unique(enty_name)
);

create table business_entity (
	buty_id SERIAL,
	buty_name VARCHAR(15),
	buty_code varchar(5),
	buty_desc varchar(55),
	buty_enty_id int,
	constraint buty_id_pk primary key(buty_id),
	constraint buty_name_uq unique(buty_name),
	constraint buty_code_uq unique(buty_code),
	constraint buty_enty_id_fk foreign key(buty_enty_id) references entity_type(enty_id)
	
);

create table users(
	user_id serial,
	user_password varchar(125),
	user_name varchar(35),
	user_email varchar(35),
	user_handphone varchar(15),
	user_create_on date,
	constraint user_id_pk primary key(user_id),
	constraint user_name_uq unique(user_name),
	constraint user_email_uq unique(user_email),
	constraint user_handphone_uq unique(user_handphone)
);

create table user_accounts(
	usac_id serial,
	usac_account_no varchar(30),
	usac_balance decimal(18,2),
	usac_create_on date,
	usac_buty_id int,
	usac_user_id int,
	constraint usac_id_pk primary key(usac_id),
	constraint usac_account_no_uq unique(usac_account_no),
	constraint usac_buty_id_fk foreign key(usac_buty_id) references business_entity(buty_id),
	constraint usac_user_id_fk foreign key(usac_user_id) references users(user_id)
);

create table transaction_type(
	traty_id serial,
	traty_name varchar(25),
	constraint traty_id_pk primary key(traty_id)
);

create table payment_transactions(
	patrx_no varchar(55),
	patrx_create_on date,
	patrx_debet decimal(18,2),
	patrx_credit decimal(18,2),
	patrx_notes varchar(125),
	patrx_acctno_from varchar(30),
	patrx_acctno_to varchar(30),
	patrx_patrx_ref varchar(55),
	patrx_traty_id int,
	constraint patrx_no_pk primary key(patrx_no),
	constraint patrx_acctno_from_fk foreign key(patrx_acctno_from) references user_accounts(usac_account_no),
	constraint patrx_acctno_to_fk foreign key(patrx_acctno_to) references user_accounts(usac_account_no),
	constraint patrx_patrx_ref_fk foreign key(patrx_patrx_ref) references payment_transactions(patrx_no),
	constraint patrx_traty_id_fk foreign key(patrx_traty_id) references transaction_type(traty_id)
);