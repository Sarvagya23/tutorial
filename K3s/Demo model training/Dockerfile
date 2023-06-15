FROM tensorflow/tensorflow:2.6.0

# Install scikit-learn
RUN pip install scikit-learn
RUN pip install pandas
RUN pip install matplotlib
RUN pip install plotly
RUN pip install seaborn

# Copy the Python script
#COPY mnist_dataset_analysis.py /app/mnist_dataset_analysis.py
COPY mnist_classification_with_lenet_5_architecture.py /app/mnist_classification_with_lenet_5_architecture.py

COPY train.csv /app/train.csv

COPY test.csv /app/test.csv

COPY sample_submission.csv /app/sample_submission.csv
