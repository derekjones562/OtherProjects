from django.shortcuts import render
from django.views import View
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
		number = int(request.GET['number'])
		original_number=number
		if number:
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
		return render(request, self.template_name, context={'original_number':original_number,'roman_numeral':self.RomanNumeral})
