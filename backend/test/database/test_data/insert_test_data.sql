INSERT INTO category(name, enabled, icon, type)
VALUES ('Sailor',
        true,
        'sailboat',
        'account');

INSERT INTO account(first_name, nickname, last_name, email, phone, balance, max_debt, category, enabled)
VALUES ('Vasco',
        'Cape Conqueror',
        'da Gama',
        'indianspice@capeofgoodhope.com',
        '+351 914 97 1498',
        33.55,
        100,
        1,
        true);

INSERT INTO account_option(account, key, value)
VALUES (1,
        'deceased',
        'true');

INSERT INTO unit(name)
VALUES ('ml');

INSERT INTO product_group(name)
VALUES ('Alcohol');

INSERT INTO product_group(name, parent)
VALUES ('Port Wine', 1);

INSERT INTO product (name, price, product_group, size, unit, tax, category)
VALUES ('Rota das Especiarias', 18.00, 1, 150, 1, 0.19, 1);

INSERT INTO product_visibility(category, product)
VALUES (1, 1);
