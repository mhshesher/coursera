# Assignment 3 - More Pandas
This assignment requires more individual learning then the last one did - you are encouraged to check out the pandas documentation to find functions or methods you might not have used yet, or ask questions on Stack Overflow and tag them as pandas and python related. And of course, the discussion forums are open for interaction with your peers and the course staff.<br><br><br>

### Question 1 (20%)
Load the energy data from the file <b>Energy Indicators.xls</b>, which is a list of indicators of energy supply and renewable electricity production from the United Nations for the year 2013, and should be put into a DataFrame with the variable name of <b>energy</b>.<br><br>

Keep in mind that this is an Excel file, and not a comma separated values file. Also, make sure to exclude the footer and header information from the datafile. The first two columns are unneccessary, so you should get rid of them, and you should change the column labels so that the columns are:<br><br>

['Country', 'Energy Supply', 'Energy Supply per Capita', '% Renewable']<br><br>

Convert Energy Supply to gigajoules (there are 1,000,000 gigajoules in a petajoule). For all countries which have missing data (e.g. data with "...") make sure this is reflected as np.NaN values.<br><br>

Rename the following list of countries (for use in later questions):<br><br>

"Republic of Korea": "South Korea",<br>
"United States of America": "United States",<br>
"United Kingdom of Great Britain and Northern Ireland": "United Kingdom",<br>
"China, Hong Kong Special Administrative Region": "Hong Kong"<br><br>

There are also several countries with numbers and/or parenthesis in their name. Be sure to remove these,<br>
e.g.<br><br>

'Bolivia (Plurinational State of)' should be 'Bolivia',<br>
'Switzerland17' should be 'Switzerland'.<br><br><br>

Next, load the GDP data from the file <b>world_bank.csv</b>, which is a csv containing countries' GDP from 1960 to 2015 from World Bank. Call this DataFrame <b>GDP</b>.<br><br>

Make sure to skip the header, and rename the following list of countries:<br><br>

"Korea, Rep.": "South Korea", <br>
"Iran, Islamic Rep.": "Iran",<br>
"Hong Kong SAR, China": "Hong Kong"<br><br><br>


Finally, load the Sciamgo Journal and Country Rank data for Energy Engineering and Power Technology from the file <b>scimagojr-3.xlsx</b>, which ranks countries based on their journal contributions in the aforementioned area. Call this DataFrame <b>ScimEn</b>.<br><br>

Join the three datasets: GDP, Energy, and ScimEn into a new dataset (using the intersection of country names). Use only the last 10 years (2006-2015) of GDP data and only the top 15 countries by Scimagojr 'Rank' (Rank 1 through 15).<br><br>

The index of this DataFrame should be the name of the country, and the columns should be ['Rank', 'Documents', 'Citable documents', 'Citations', 'Self-citations', 'Citations per document', 'H index', 'Energy Supply', 'Energy Supply per Capita', '% Renewable', '2006', '2007', '2008', '2009', '2010', '2011', '2012', '2013', '2014', '2015'].<br><br>

<i>This function should return a DataFrame with 20 columns and 15 entries.</i>

### Question 2 (6.6%)

The previous question joined three datasets then reduced this to just the top 15 entries. When you joined the datasets, but before you reduced this to the top 15 items, how many entries did you lose?<br><br>

<i>This function should return a single number.</i>

## Answer the following questions in the context of only the top 15 countries by Scimagojr Rank (aka the DataFrame returned by answer_one())

### Question 3 (6.6%)

What is the average GDP over the last 10 years for each country? (exclude missing values from this calculation).<br><br>

<i>This function should return a Series named avgGDP with 15 countries and their average GDP sorted in descending order.
</i>

###Question 4 (6.6%)

By how much had the GDP changed over the 10 year span for the country with the 6th largest average GDP?<br><br>

<i>This function should return a single number.</i>

### Question 5 (6.6%)

What is the mean Energy Supply per Capita?<br><br>

<i>This function should return a single number.</i>

### Question 6 (6.6%)

What country has the maximum % Renewable and what is the percentage?<br><br>

<i>This function should return a tuple with the name of the country and the percentage.</i>

### Question 7 (6.6%)

Create a new column that is the ratio of Self-Citations to Total Citations. What is the maximum value for this new column, and what country has the highest ratio?<br><br>

<i>This function should return a tuple with the name of the country and the ratio.</i>

### Question 8 (6.6%)

Create a column that estimates the population using Energy Supply and Energy Supply per capita. What is the third most populous country according to this estimate?<br><br>

<i>This function should return a single string value.</i>

### Question 9 (6.6%)

Create a column that estimates the number of citable documents per person. What is the correlation between the number of citable documents per capita and the energy supply per capita? Use the .corr() method, (Pearson's correlation).

<i>This function should return a single number.

(Optional: Use the built-in function plot9() to visualize the relationship between Energy Supply per Capita vs. Citable docs per Capita)</i>

### Question 10 (6.6%)

Create a new column with a 1 if the country's % Renewable value is at or above the median for all countries in the top 15, and a 0 if the country's % Renewable value is below the median.<br><br>

<i>This function should return a series named HighRenew whose index is the country name sorted in ascending order of rank.</i>

### Question 11 (6.6%)

Use the following dictionary to group the Countries by Continent, then create a dateframe that displays the sample size (the number of countries in each continent bin), and the sum, mean, and std deviation for the estimated population of each country.<br><br>

ContinentDict  = {'China':'Asia', <br>
                  'United States':'North America', <br>
                  'Japan':'Asia', <br>
                  'United Kingdom':'Europe', <br>
                  'Russian Federation':'Europe', <br>
                  'Canada':'North America', <br>
                  'Germany':'Europe', <br>
                  'India':'Asia',<br>
                  'France':'Europe', <br>
                  'South Korea':'Asia', <br>
                  'Italy':'Europe', <br>
                  'Spain':'Europe', <br>
                  'Iran':'Asia',<br>
                  'Australia':'Australia', <br>
                  'Brazil':'South America'}<br><br>

<i>This function should return a DataFrame with index named Continent ['Asia', 'Australia', 'Europe', 'North America', 'South America'] and columns ['size', 'sum', 'mean', 'std']</i>

### Question 12 (6.6%)

Cut % Renewable into 5 bins. Group Top15 by the Continent, as well as these new % Renewable bins. How many countries are in each of these groups?<br><br>

<i>This function should return a <b>Series</b> with a MultiIndex of Continent, then the bins for % Renewable. Do not include groups with no countries.</i>

### Question 13 (6.6%)

Convert the Population Estimate series to a string with thousands separator (using commas). Do not round the results.<br><br>

e.g. 317615384.61538464 -> 317,615,384.61538464<br><br>

<i>This function should return a Series PopEst whose index is the country name and whose values are the population estimate string.
</i>

