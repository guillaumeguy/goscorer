install.packages("pmml")
install.packages('randomForest')

library("XML")
library("pmml")
library('randomForest')

names(iris) = c('sepal_length_cm', 'sepal_width_cm', 'petal_length_cm', 'petal_width_cm','species')

iris$is_setosa <- as.factor(ifelse(iris$species == 'setosa',1,0))

iris_rf <- randomForest(is_setosa ~ sepal_length_cm + sepal_width_cm + petal_length_cm + petal_width_cm,
                           data = iris,
                          ntree = 20
                           )
pmml_obj <- pmml(iris_rf)

saveXML(pmml_obj, "~/go/src/goscorer/static/iris.pmml")