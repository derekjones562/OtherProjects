from __future__ import unicode_literals

import uuid

from datetime import date
from django.contrib.auth.models import User
from django.db import models

# Create your models here.
from rest_framework import serializers


class WordSmith(models.Model):

    first_name = models.CharField(max_length=50)
    last_name = models.CharField(max_length=50)
    user = models.OneToOneField(User, primary_key=True, unique=True)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def __str__(self):
        return "Word Smith: {} {}".format(self.first_name, self.last_name)


class Tome(models.Model):

    word_smith = models.ForeignKey(WordSmith, related_name='tomes', on_delete=models.CASCADE)
    tome_title = models.CharField(max_length=255)
    number_of_pages = models.IntegerField(default=0)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def __str__(self):
        return "{}: {}".format("Tome: ", self.tome_title)


class Page(models.Model):

    page_id = models.CharField(max_length=255, unique=True, default=uuid.uuid4)
    page_date = models.DateField(default=date.today())
    tome_id = models.ForeignKey(Tome, related_name='pages', on_delete=models.CASCADE)
    entry = models.TextField()

    def save(self, force_insert=False, force_update=False, using=None, update_fields=None):
        super().save(force_insert, force_update, using, update_fields)
        print(Page.objects.get(page_id=self.page_id))
        if Page.objects.get(page_id=self.page_id):
            tome = Tome.objects.get(id=self.tome_id.id)
            tome.number_of_pages += 1
            tome.save()

    def delete(self, using=None, keep_parents=False):
        super().delete(using=using, keep_parents=keep_parents)
        tome = Tome.objects.get(id=self.tome_id.id)
        tome.number_of_pages -= 1
        tome.save()

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def __str__(self):
        return "Page: {} --- {}".format(self.id, self.page_date)
