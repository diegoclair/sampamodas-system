package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
//Only work doing 1 command per version, you cannot create two tables in the same script, need to create a new version
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Create tab_company",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_company (
					id INT NOT NULL AUTO_INCREMENT,
					document_number VARCHAR(14) NOT NULL,
					legal_name VARCHAR(45) NOT NULL,
					commercial_name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     2,
			Description: "Create tab_business",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_business (
					id INT NOT NULL AUTO_INCREMENT,
					company_id INT NOT NULL,
					name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
					INDEX fk_tab_business_tab_company1_idx (company_id ASC) VISIBLE,
					CONSTRAINT fk_tab_business_tab_company1
					FOREIGN KEY (company_id)
					REFERENCES tab_company (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     3,
			Description: "Create tab_brand",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_brand (
					id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NOT NULL,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     4,
			Description: "Create tab_product",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_product (
					id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NOT NULL,
					cost DECIMAL NULL DEFAULT 0.00,
					price DECIMAL NULL DEFAULT 0.00,
					gender_type VARCHAR(45) NOT NULL,
					business_id INT NOT NULL,
					brand_id INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
					INDEX fk_tab_products_tab_company1_idx (business_id ASC) VISIBLE,
					INDEX fk_tab_product_tab_brand1_idx (brand_id ASC) VISIBLE,
					CONSTRAINT fk_tab_products_tab_company1
					FOREIGN KEY (business_id)
					REFERENCES tab_business (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_product_tab_brand1
					FOREIGN KEY (brand_id)
					REFERENCES tab_brand (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE = InnoDB;
				`,
		},
		{
			Version:     5,
			Description: "Create tab_lead",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_lead (
					id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NULL,
					email VARCHAR(45) NULL,
					phone_number VARCHAR(45) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     6,
			Description: "Create tab_send_method",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_send_method (
					id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			  
			`,
		},
		{
			Version:     7,
			Description: "Create tab_payment_method",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_payment_method (
					id INT NOT NULL AUTO_INCREMENT,
					method_name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     8,
			Description: "Create tab_sale",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_sale (
					id INT NOT NULL AUTO_INCREMENT,
					lead_id INT NOT NULL,
					total_price DECIMAL NOT NULL DEFAULT 0.00,
					freight DECIMAL GENERATED ALWAYS AS (0.00) VIRTUAL,
					payment_method_id INT NULL,
					send_method_id INT NULL,
					PRIMARY KEY (id),
					INDEX fk_tab_sale_tab_lead1_idx (lead_id ASC) VISIBLE,
					INDEX fk_tab_sale_tab_send_method1_idx (send_method_id ASC) VISIBLE,
					INDEX fk_tab_sale_tab_payment_method1_idx (payment_method_id ASC) VISIBLE,
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
					CONSTRAINT fk_tab_sale_tab_lead1
					FOREIGN KEY (lead_id)
					REFERENCES tab_lead (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_sale_tab_send_method1
					FOREIGN KEY (send_method_id)
					REFERENCES tab_send_method (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_sale_tab_payment_method1
					FOREIGN KEY (payment_method_id)
					REFERENCES tab_payment_method (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     9,
			Description: "Create tab_product_stock",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_product_stock (
					id INT NOT NULL,
					product_id INT NOT NULL,
					color VARCHAR(45) NULL,
					size VARCHAR(45) NOT NULL,
					quantity INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					INDEX fk_tab_product_quantity_tab_products1_idx (product_id ASC) VISIBLE,
					CONSTRAINT fk_tab_product_quantity_tab_products1
					FOREIGN KEY (product_id)
					REFERENCES tab_product (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     10,
			Description: "Create tab_lead_address",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_lead_address (
					id INT NOT NULL AUTO_INCREMENT,
					lead_id INT NOT NULL,
					address_type VARCHAR(45) NULL COMMENT 'if its home, work, etc..',
					street VARCHAR(100) NULL,
					number VARCHAR(45) NULL,
					neighborhood VARCHAR(45) NULL,
					complement VARCHAR(45) NULL,
					city VARCHAR(45) NULL,
					federative_unit VARCHAR(2) NULL,
					zip_code VARCHAR(45) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					INDEX fk_tab_lead_address_tab_lead1_idx (lead_id ASC) INVISIBLE,
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
					CONSTRAINT fk_tab_lead_address_tab_lead1
					FOREIGN KEY (lead_id)
					REFERENCES tab_lead (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     11,
			Description: "Create tab_sale_product",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_sale_product (
					id INT NOT NULL AUTO_INCREMENT,
					sale_id INT NOT NULL,
					product_id INT NOT NULL,
					quantity INT NOT NULL,
					price DECIMAL NOT NULL DEFAULT 0.00,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (id),
					UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
					INDEX fk_tab_sale_products_tab_sale1_idx (sale_id ASC) VISIBLE,
					INDEX fk_tab_sale_products_tab_products1_idx (product_id ASC) VISIBLE,
					CONSTRAINT fk_tab_sale_products_tab_sale1
					FOREIGN KEY (sale_id)
					REFERENCES tab_sale (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_sale_products_tab_products1
					FOREIGN KEY (product_id)
					REFERENCES tab_product (id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
	}
)
