\c avhbs;

INSERT INTO category(name, enabled, icon, type)
VALUES ('Active B',
        true,
        'icon',
        'account');

INSERT INTO account(first_name, nickname, last_name, email, phone, balance, max_debt, category, enabled)
VALUES ('Brandon',
        'Broseidon',
        'Poseidon', 'kingofocean@mail.com',
        '+30 123 456 78',
        33.55,
        100,
        1,
        true);

INSERT INTO account_option(account, key, value)
VALUES (1,
        'opt1',
        'val1');

INSERT INTO unit(name)
VALUES('ml');