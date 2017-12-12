CREATE OR REPLACE FUNCTION add_currency(name VARCHAR(200), logo_url CHAR(500)) 
RETURNS void AS $$
BEGIN
  INSERT INTO "currency" VALUES (DEFAULT, name, logo_url);
END;
$$ LANGUAGE plpgsql;

SELECT add_currency('Bitcoin', 'https://bitcoin.org/img/icons/opengraph.png');
SELECT add_currency('Ethereum', 'http://files.coinmarketcap.com.s3-website-us-east-1.amazonaws.com/static/img/coins/200x200/ethereum.png');
SELECT add_currency('Litecoin', 'https://litecoin.org/img/litecoin.svg');

CREATE OR REPLACE FUNCTION transactions_by_date(date_start TIMESTAMP, date_end TIMESTAMP)
DECLARE ref refcursor;
BEGIN
  OPEN ref FOR SELECT * FROM "transaction" WHERE date >= date_start AND date < date_end;
  RETURN ref;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION user_investment(in_user_id INT)
DECLARE ref refcursor;
BEGIN
  OPEN ref FOR SELECT SUM(price) AS investment
               FROM "transaction" t, "balance" b
               WHERE b.user_id=in_user_id AND b.id=t.balance_id;
  RETURN ref;
END;
$$ LANGUAGE plpgsql;
