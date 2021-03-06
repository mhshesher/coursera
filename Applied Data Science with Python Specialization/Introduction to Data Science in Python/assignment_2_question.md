# Assignment 2 - Pandas Introduction
All questions are weighted the same in this assignment.<br><br>


## Part 1

The following code loads the olympics dataset (olympics.csv), which was derrived from the Wikipedia entry on [All Time Olympic Games Medals](https://en.wikipedia.org/wiki/All-time_Olympic_Games_medal_table), and does some basic data cleaning.<br><br>
The columns are organized as # of Summer games, Summer medals, # of Winter games, Winter medals, total # number of games, total # of medals. Use this dataset to answer the questions below.<br><br>

### Question 0 (Example)
What is the first country in df?<br><br>
<i>This function should return a Series.</i>

### Question 1
Which country has won the most gold medals in summer games?<br><br>
<i>This function should return a single string value.</i>

### Question 2
Which country had the biggest difference between their summer and winter gold medal counts?<br><br>
<i>This function should return a single string value.</i>

### Question 3
Which country has the biggest difference between their summer gold medal counts and winter gold medal counts relative to their total gold medal count?

(Summer Gold−Winter Gold) / Total Gold

Only include countries that have won at least 1 gold in both summer and winter.<br><br>
<i>This function should return a single string value.</i>

### Question 4
Write a function that creates a Series called "Points" which is a weighted value where each gold medal (Gold.2) counts for 3 points, silver medals (Silver.2) for 2 points, and bronze medals (Bronze.2) for 1 point. The function should return only the column (a Series object) which you created, with the country names as indices.<br><br>
<i>This function should return a Series named Points of length 146</i>

## Part 2
For the next set of questions, we will be using census data from the [United States Census Bureau](https://www.census.gov/). Counties are political and geographic subdivisions of states in the United States. This dataset contains population data for counties and states in the US from 2010 to 2015.<br><br>
The census dataset (census.csv) should be loaded as census_df. Answer questions using this as appropriate.<br><br>

### Question 5
Which state has the most counties in it? (hint: consider the sumlevel key carefully! You'll need this for future questions too...)<br><br>
<i>This function should return a single string value.</i>

### Question 6
<b>Only looking at the three most populous counties for each state</b>, what are the three most populous states (in order of highest population to lowest population)? Use CENSUS2010POP.<br><br>
<i>This function should return a list of string values.</i>

### Question 7
Which county has had the largest absolute change in population within the period 2010-2015? (Hint: population values are stored in columns POPESTIMATE2010 through POPESTIMATE2015, you need to consider all six columns.)<br><br>
e.g. If County Population in the 5 year period is 100, 120, 80, 105, 100, 130, then its largest change in the period would be |130-80| = 50.<br><br>
<i>This function should return a single string value.</i>

### Question 8
In this datafile, the United States is broken up into four regions using the "REGION" column.<br><br>
Create a query that finds the counties that belong to regions 1 or 2, whose name starts with 'Washington', and whose POPESTIMATE2015 was greater than their POPESTIMATE 2014.<br><br>
<i>This function should return a 5x2 DataFrame with the columns = ['STNAME', 'CTYNAME'] and the same index ID as the census_df (sorted ascending by index).</i>
