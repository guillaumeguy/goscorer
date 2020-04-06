install.packages("pmml")
install.packages('randomForest')

library("XML")
library("pmml")
library('randomForest')

names(iris) = c('sepal_length_cm', 'sepal_width_cm', 'petal_length_cm', 'petal_width_cm','species')


saveXML(pmml_obj, "~/go/src/goscorer/static/iris.pmml")
