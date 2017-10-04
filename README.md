# chalenge-ha (HousingAnywhere chalenge)
Just a chalenge

```
Usage of ./Chalenge: 
  -csv string                                         
        Inform a valid csv file to be processed. (default "./geoData.csv")
  -ds int
        Configure the source of the data. you can inform 0 for csv or 1 for PostgreSQL (under development)
  -lg float
        Inform the reference longitude to measure the distance. (default 4.478617)
  -lt float
        Inform the reference latitude to measure the distance. (default 51.925146)
  -n int
        Inform the max relevant elements you want. (default 5)
```

Here you can see the times that i achieve with this program:

<p align="center">
  <img src="https://github.com/softctrl/chalenge-ha/blob/master/time00.png?raw=true" alt="Proccess geoData.csv"/>
</p>

<p align="center">
  <img src="https://github.com/softctrl/chalenge-ha/blob/master/time01.png?raw=true" alt="Proccess 1M.csv"/>
</p>

<p align="center">
  <img src="https://github.com/softctrl/chalenge-ha/blob/master/time02.png?raw=true" alt="Proccess 50M.csv"/>
</p>

# Simulating a large Dataset (how i test with 50M of records!!)

Just download the [scgen.py](https://github.com/softctrl/chalenge-ha/blob/master/scgen.py) file and execute:

```
python3.6 scgen.py > geo_data_set.csv
```

You can change the MAX variable on this program to a value that you need to create a test csv file with location values. Also i am using Python 3.6.
