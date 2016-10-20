from django.shortcuts import render
from django.views import View
from django.shortcuts import Http404
from .models import Conversion
# Create your views here.

class ConverterView(View):
	template_name = 'converter.html'
	RomanNumeral = ''
	
	def add_letter(self, number,increment, letter):
		if(number%increment is not number):
			while(number>increment-1):
				self.RomanNumeral = self.RomanNumeral + letter
				number = number -increment
		return number
		
		
	def get(self, request):
		url_num = request.GET.get('number')
		page = render(request, self.template_name)
		if url_num:
			number = int(url_num) 
			original_number=number
			existing_in_db = Conversion.objects.all().filter(starting_num=original_number)
			if existing_in_db.exists():
				previous_times_converted=Conversion.objects.get(starting_num=original_number).times_converted
				Conversion.objects.select_related().filter(starting_num=original_number).update(times_converted= previous_times_converted + 1)
				self.RomanNumeral=existing_in_db.values('roman_numeral')[0].get('roman_numeral')
			else:
				number = self.add_letter(number,1000, 'M')
				number = self.add_letter(number,900, 'CM')
				number = self.add_letter(number,500, 'D')
				number = self.add_letter(number,400, 'CD')
				number = self.add_letter(number,100, 'C')
				number = self.add_letter(number,90, 'XC')
				number = self.add_letter(number,50, 'L')
				number = self.add_letter(number,40, 'XL')
				number = self.add_letter(number,10, 'X')
				number = self.add_letter(number,9, 'IX')
				number = self.add_letter(number,5, 'V')
				number = self.add_letter(number,4, 'IV')
				number = self.add_letter(number,1, 'I')	
				new_convert = Conversion(roman_numeral= self.RomanNumeral, starting_num=original_number)
				new_convert.save()
			page = render(request, self.template_name, context={'original_number':original_number,'roman_numeral':self.RomanNumeral})
		return page