from __future__ import unicode_literals

from django.db import models

# Create your models here.
class Conversion(models.Model):
	roman_numeral = models.CharField(max_length=250)
	starting_num = models.IntegerField()
	times_converted = models.IntegerField(default=1)
	
	def __str__(self):
		return str(self.starting_num) +" to "+self.roman_numeral 