from django.conf.urls import url
from rest_framework.urlpatterns import format_suffix_patterns
from apps.tome.views import ui, api

urlpatterns = [
               # url(r'^pages/$', api.page_list), # FBV
               # url(r'^pages/(?P<pk>[0-9]+)/$', api.page_detail), #FBV
               url(r'^word_smiths/$', api.WordSmithList.as_view(), name='wordsmith-list'),
               url(r'^word_smiths/(?P<pk>[0-9]+)/$', api.WordSmithDetail.as_view(), name='wordsmith-detail'),
               url(r'^tomes/$', api.TomeList.as_view(), name='tome-list'),
               url(r'^tomes/(?P<pk>[0-9]+)/$', api.TomeDetail.as_view(), name='tome-detail'),
               url(r'^pages/$', api.PageList.as_view(), name='page-list'),
               url(r'^pages/(?P<pk>[0-9]+)/$', api.PageDetail.as_view(), name='page-detail'),



               url(r'login', ui.login, name='login'),
               url(r'create_wordsmith', ui.create_wordsmith, name='create_wordsmith'),
               url(r'', ui.my_tomes, name='my_tomes'),
]

urlpatterns = format_suffix_patterns(urlpatterns)