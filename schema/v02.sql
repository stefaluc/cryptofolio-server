CREATE OR REPLACE FUNCTION add_currency(name VARCHAR(200), logo_url CHAR(500)) 
RETURNS void AS $$
BEGIN
  INSERT INTO "currency" VALUES (DEFAULT, name, logo_url);
END;
$$ LANGUAGE plpgsql;

SELECT add_currency('Bitcoin', 'https://bitcoin.org/img/icons/opengraph.png');
SELECT add_currency('Ethereum', 'http://files.coinmarketcap.com.s3-website-us-east-1.amazonaws.com/static/img/coins/200x200/ethereum.png');
SELECT add_currency('Litecoin', 'https://litecoin.org/img/litecoin.svg');
