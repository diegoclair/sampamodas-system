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
					company_id INT NOT NULL AUTO_INCREMENT,
					document_number VARCHAR(14) NOT NULL,
					legal_name VARCHAR(45) NOT NULL,
					commercial_name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (company_id),
					UNIQUE INDEX id_UNIQUE (company_id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     2,
			Description: "Create tab_business",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_business (
					business_id INT NOT NULL AUTO_INCREMENT,
					company_id INT NOT NULL,
					name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (business_id),
					UNIQUE INDEX id_UNIQUE (business_id ASC) VISIBLE,
					INDEX fk_tab_business_tab_company1_idx (company_id ASC) VISIBLE,
					CONSTRAINT fk_tab_business_tab_company1
					FOREIGN KEY (company_id)
					REFERENCES tab_company (company_id)
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
					brand_id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (brand_id),
					UNIQUE INDEX id_UNIQUE (brand_id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     4,
			Description: "Create tab_gender",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_gender (
					gender_id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (gender_id),
					UNIQUE INDEX id_UNIQUE (gender_id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     5,
			Description: "Create tab_supplier",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_supplier (
					supplier_id INT NOT NULL AUTO_INCREMENT,
					business_id INT NOT NULL,
					name VARCHAR(45) NULL,
					email VARCHAR(45) NULL,
					phone_number VARCHAR(45) NOT NULL,
					instagram VARCHAR(45) NULL,
					website VARCHAR(85) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (supplier_id),
					UNIQUE INDEX supplier_id_UNIQUE (supplier_id ASC) VISIBLE,
					INDEX fk_tab_supplier_tab_business1_idx (business_id ASC) VISIBLE,
					CONSTRAINT fk_tab_supplier_tab_business1
					FOREIGN KEY (business_id)
					REFERENCES tab_business (business_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     6,
			Description: "Create tab_product",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_product (
					product_id INT NOT NULL AUTO_INCREMENT,
					supplier_id INT NOT NULL,
					name VARCHAR(45) NOT NULL,
					cost DECIMAL(7,2) NULL DEFAULT 0.00,
					price DECIMAL(7,2) NULL DEFAULT 0.00,
					gender_id INT NOT NULL,
					brand_id INT NOT NULL,
					business_id INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (product_id),
					UNIQUE INDEX id_UNIQUE (product_id ASC) VISIBLE,
					INDEX fk_tab_products_tab_business1_idx (business_id ASC) VISIBLE,
					INDEX fk_tab_product_tab_brand1_idx (brand_id ASC) VISIBLE,
					INDEX fk_tab_product_tab_gender1_idx (gender_id ASC) VISIBLE,
					INDEX fk_tab_product_tab_supplier1_idx (supplier_id ASC) VISIBLE,
					CONSTRAINT fk_tab_product_tab_business1
					FOREIGN KEY (business_id)
					REFERENCES tab_business (business_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_product_tab_brand1
					FOREIGN KEY (brand_id)
					REFERENCES tab_brand (brand_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_product_tab_gender1
					FOREIGN KEY (gender_id)
					REFERENCES tab_gender (gender_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_product_tab_supplier1
					FOREIGN KEY (supplier_id)
					REFERENCES tab_supplier (supplier_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
				`,
		},
		{
			Version:     7,
			Description: "Create tab_lead",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_lead (
					lead_id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NULL,
					email VARCHAR(45) NULL,
					phone_number VARCHAR(45) NOT NULL,
					instagram VARCHAR(45) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (lead_id, phone_number),
					UNIQUE INDEX id_UNIQUE (lead_id ASC) VISIBLE,
					UNIQUE INDEX phone_number_UNIQUE (phone_number ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     8,
			Description: "Create tab_send_method",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_send_method (
					send_method_id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (send_method_id),
					UNIQUE INDEX id_UNIQUE (send_method_id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			  
			`,
		},
		{
			Version:     9,
			Description: "Create tab_payment_method",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_payment_method (
					payment_method_id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (payment_method_id),
					UNIQUE INDEX id_UNIQUE (payment_method_id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     10,
			Description: "Create tab_sale",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_sale (
					sale_id INT NOT NULL AUTO_INCREMENT,
					lead_id INT NOT NULL,
					total_price DECIMAL(7,2) NOT NULL DEFAULT 0.00,
					freight DECIMAL(7,2) NOT NULL,
					payment_method_id INT NOT NULL,
					send_method_id INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (sale_id),
					INDEX fk_tab_sale_tab_lead1_idx (lead_id ASC) VISIBLE,
					INDEX fk_tab_sale_tab_send_method1_idx (send_method_id ASC) VISIBLE,
					INDEX fk_tab_sale_tab_payment_method1_idx (payment_method_id ASC) VISIBLE,
					UNIQUE INDEX id_UNIQUE (sale_id ASC) VISIBLE,
					CONSTRAINT fk_tab_sale_tab_lead1
					FOREIGN KEY (lead_id)
					REFERENCES tab_lead (lead_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_sale_tab_send_method1
					FOREIGN KEY (send_method_id)
					REFERENCES tab_send_method (send_method_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_sale_tab_payment_method1
					FOREIGN KEY (payment_method_id)
					REFERENCES tab_payment_method (payment_method_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     11,
			Description: "Create tab_color",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_color (
					color_id INT NOT NULL AUTO_INCREMENT,
					name VARCHAR(45) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (color_id),
					UNIQUE INDEX id_UNIQUE (color_id ASC) VISIBLE)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     12,
			Description: "Create tab_product_stock",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_product_stock (
					product_stock_id INT NOT NULL AUTO_INCREMENT,
					product_id INT NOT NULL,
					color_id INT NOT NULL,
					size VARCHAR(45) NOT NULL,
					available_quantity INT NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (product_stock_id),
					INDEX fk_tab_product_quantity_tab_products1_idx (product_id ASC) VISIBLE,
					INDEX fk_tab_product_stock_tab_color1_idx (color_id ASC) VISIBLE,
					UNIQUE INDEX id_UNIQUE (product_stock_id ASC) VISIBLE,
					CONSTRAINT fk_tab_product_stock_tab_product1
					FOREIGN KEY (product_id)
					REFERENCES tab_product (product_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_product_stock_tab_color1
					FOREIGN KEY (color_id)
					REFERENCES tab_color (color_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     13,
			Description: "Create tab_lead_address",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_lead_address (
					lead_addressid INT NOT NULL AUTO_INCREMENT,
					lead_id INT NOT NULL,
					address_type VARCHAR(45) NULL COMMENT 'if its home, work, etc..',
					street VARCHAR(100) NOT NULL,
					number VARCHAR(45) NOT NULL,
					neighborhood VARCHAR(45) NOT NULL,
					complement VARCHAR(45) NOT NULL,
					city VARCHAR(45) NOT NULL,
					federative_unit VARCHAR(2) NOT NULL,
					zip_code VARCHAR(9) NOT NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (lead_addressid),
					INDEX fk_tab_lead_address_tab_lead1_idx (lead_id ASC) INVISIBLE,
					UNIQUE INDEX id_UNIQUE (lead_addressid ASC) VISIBLE,
					CONSTRAINT fk_tab_lead_address_tab_lead1
					FOREIGN KEY (lead_id)
					REFERENCES tab_lead (lead_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     14,
			Description: "Create tab_sale_product",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_sale_product (
					sale_product_id INT NOT NULL AUTO_INCREMENT,
					sale_id INT NOT NULL,
					product_stock_id INT NOT NULL,
					quantity INT NOT NULL,
					price DECIMAL(7,2) NOT NULL DEFAULT 0.00,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (sale_product_id),
					UNIQUE INDEX id_UNIQUE (sale_product_id ASC) VISIBLE,
					INDEX fk_tab_sale_products_tab_sale1_idx (sale_id ASC) VISIBLE,
					INDEX fk_tab_sale_product_tab_product_stock1_idx (product_stock_id ASC) VISIBLE,
					CONSTRAINT fk_tab_sale_products_tab_sale1
					FOREIGN KEY (sale_id)
					REFERENCES tab_sale (sale_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION,
					CONSTRAINT fk_tab_sale_product_tab_product_stock1
					FOREIGN KEY (product_stock_id)
					REFERENCES tab_product_stock (product_stock_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     15,
			Description: "Create tab_stock_input",
			Script: `
				CREATE TABLE IF NOT EXISTS tab_stock_input (
					stock_input_id INT UNSIGNED NOT NULL,
					product_stock_id INT NOT NULL,
					quantity VARCHAR(45) NULL,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
					PRIMARY KEY (stock_input_id),
					UNIQUE INDEX id_UNIQUE (stock_input_id ASC) VISIBLE,
					INDEX fk_tab_stock_tab_product_stock1_idx (product_stock_id ASC) VISIBLE,
					CONSTRAINT fk_tab_stock_tab_product_stock1
					FOREIGN KEY (product_stock_id)
					REFERENCES tab_product_stock (product_stock_id)
					ON DELETE NO ACTION
					ON UPDATE NO ACTION)
				ENGINE=InnoDB CHARACTER SET=utf8;
			`,
		},
		{
			Version:     16,
			Description: "Insert data into tab_payment_method",
			Script: `
				INSERT INTO tab_payment_method 
					(payment_method_id, name) VALUES
					(1, "Cartão de Débito"),
					(2, "Cartão de Crédito"),
					(3, "PIX/Transferência"),
					(4, "Dinheiro");
			`,
		},
		{
			Version:     17,
			Description: "Insert data into tab_send_method",
			Script: `
				INSERT INTO tab_send_method 
					(send_method_id, name) VALUES
					(1, "Correios"),
					(2, "Em mãos"),
					(3, "Motoboy"),
					(4, "Loggi");
			`,
		},
	}
)
