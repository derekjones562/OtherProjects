from django.conf.urls import url
from . import views
app_name = 'romannumeralconverter'
urlpatterns = [
    url(r'^$', views.ConverterView.as_view(), name='converter' ),
    #url(r'^(?P<number>\d{1,5})$', views.DisplayConversionView.as_view(), name='display'),
]