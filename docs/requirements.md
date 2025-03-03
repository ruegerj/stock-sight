# Stories & Requirements

## Scope

Stock splits (price to high, one share gets split into X) are considered out of scope and have to be managed by hand. Short sells
(selling more stocks than you own) and leveraged buys (buying stocks with lent money) aren't supported.

## Requirements

**Functional:**

1. I want to select stocks either traded on _bxswiss_ and or _Six_
2. I can choose some of the selectable stocks in order to track their price development from this point forward.
3. I can track the point of time when I sold/bought a concrete amount of stocks and their current value at this time.
4. I can visualize the price history of a tracked stock over the time periods: day, week, month & year.
5. I can visualize the buying/selling points of my stocks within the price chart.
6. I can visualize my current stock portfolio, based on the buys/sells previously registered in the app.

**Non-Functional:**

1. The app should start in under 5sec.
2. The app should run on Linux (AMD64) / Windows (AMD64) / Mac (ARM64)
3. The app is designed in an extensible way in order to facilitate adapting new marketplaces & item types.
4. The app and its datastore is portable and can be moved to another device.
5. The app's data is encrypted at rest (Optional).
6. The app should display the portfolio (based on latest data points) and price-history of the tracked stocks offline.

## Stories

### Story-1: Selecting Stocks (_Must_)

As a user I want to be able to select a stock, whose future price development I'd like to track with the application.

**Acceptance criteria:**

- When a ticker (unique stock identifier) is entered the stock is added to the tracked stocks of the user

### Story-2: Display of tracked Stocks (_Must_)

As a user I want to be able to display all the stocks which I'm currently tracking with the application, in order to gain an overview.

**Acceptance criteria:**

- All currently tracked stocks are displayed with their:
  - ticker
  - full name
  - current price & currency
- The current price should be sourced from the latest data point of the given stock.
- The price should be displayed with the timestamp of the datapoint.

### Story-3: Adding stock transactions (_Should_)

As a user I want to be able to register a stock transaction, in order to save create a buy/sell history.

**Acceptance criteria:**

- Transaction can only be created for previously _tracked_ stocks
- A transaction must inquire and thus contain the following data:
  - ticker
  - amount
  - timestamp
  - transaction kind (buy/sell)
  - current price per unit (including currency)
- The quantity of "sell"-transactions can be at most the sum of all amounts of "buy"-transaction of this stock
- The transaction date can be the current timestamp or one in the past, but not in the future

### Story-4: Price Histogram (_Must_)

As a user I want to be able to see the price history of a selected stock, visualized as a line-chart.

**Acceptance criteria:**

- The chart can only be displayed for previously tracked stocks
- The following timespans can be displayed:
  - last day
  - last week
  - last month
  - last year
  - year to date (from 01.01.xxxx to now)
- If the collected data-points don't cover the selected timespan, the chart is rendered according to the "best-effort" principle.

### Story-5: Price Statistics for visualized Timespan (_Could_)

As a user I want to be able to view the price statistics: _min_, _max_, _avg_, _mean_, _performance in %_ and _absolute performance_ for the currently
selected timestan of the price-chart.

**Acceptance criteria:**

- The required price statistics are displayed for the selected timespan
- The amount of data points is displayed besides the statistics

### Story-6: Display of Transactions in Chart (_Should_)

As a user I want to see my buys & sells in the price histogram, in order to gain insights on my decisions.

**Acceptance criteria:**

- All registered transactions are displayed in the chart
- For each transaction, it's type (buy/sell) is displayed visually distinct from the others

### Story-7: Stock Portfolio (_Should_)

As a user I want to see my current stock portfolio based my previously registered transactions, in order to have an overview over my assets.

**Acceptance criteria:**

- A stock should be displayed if the users currently holds at least some quantity of a share (> 0)
- For every held stock the following information should be displayed:
  - ticker
  - name
  - amount of shares currently held
  - current price (according to latest data point for stock)
