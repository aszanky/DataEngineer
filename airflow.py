from datetime import datetime, timedelta
from airflow import DAG
from airflow.models import Variable
from airflow.operators.python_operator import PythonOperator

import insertProduct

default_args = {
    'owner': 'kitabisa',
    'depends_on_past': False,
    'start_date': datetime(2020, 7, 12),
    'email': ['kitabisa@kitabisa.co.id'],
    'email_on_failure': False,
    'email_on_retry': False,
    'retries': 1,
    'retry_delay': timedelta(minutes=5),
}

dag = DAG(
    'dag_insert_product',
    default_args=default_args,
    schedule_interval="*/1 * * * *"
)

process_dag = PythonOperator(
    task_id='dag_insert_product',
    python_callable=insertProduct,
    dag=dag
)