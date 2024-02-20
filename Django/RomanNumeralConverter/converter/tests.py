from django.test import TestCase
from converter.models import Conversion
from converter.views import ConverterView

# Create your tests here.
class ConverterTestCase(TestCase):
		
	def test_base_url(self):
		resp = self.client.get('/')
		self.assertEqual(resp.status_code, 200)
		self.assertEqual(resp.resolver_match.func.__name__, ConverterView.as_view().__name__)
		
		
	def test_add_letter(self):
		#add_letter(self, number,increment, letter)
		add_tester = ConverterView()
		test_number = add_tester.add_letter(1200,1000, 'F')
		self.assertEqual(test_number,200)
		test_number = add_tester.add_letter(800,400, 'O')
		self.assertEqual(test_number,0)
		test_number = add_tester.add_letter(401,400, 'B')
		self.assertEqual(test_number,1)
		test_number = add_tester.add_letter(799,400, 'A')
		self.assertEqual(test_number,399)
		test_number = add_tester.add_letter(3,2, 'R')
		self.assertEqual(test_number,1)
		self.assertEqual(add_tester.RomanNumeral,'FOOBAR')
		
		
	def test_get(self):
		#get(self, request)
		
		'''conversion_test = Conversion.objects.create(
			roman_numeral = 'MCMXLI',
			starting_num = 1941,
			times_converted = 1,
		)'''
		resp = self.client.get('/romannumeralconverter/')
		self.assertEqual(resp.status_code, 200)
		self.assertEqual(resp.templates[0].name, 'converter.html')
		
	
		resp_with_number = self.client.get('/romannumeralconverter/?number=1941')
		self.assertEqual(resp_with_number.status_code, 200)
		self.assertEqual(resp_with_number.templates[0].name, 'converter.html')
		self.assertEqual(resp_with_number.context['original_number'],1941)
		self.assertEqual(resp_with_number.context['roman_numeral'],'MCMXLI')