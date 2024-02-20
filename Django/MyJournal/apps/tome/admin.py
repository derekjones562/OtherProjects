from django.contrib import admin

# Register your models here.
from apps.tome.models import *

admin.site.register(WordSmith)
admin.site.register(Tome)
admin.site.register(Page)
