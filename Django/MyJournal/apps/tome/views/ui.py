
from django.shortcuts import render_to_response
from django.template import RequestContext

from apps.tome.controllers import TomeController
from apps.tome.models import *

def login(request):
    if request.method == 'POST':
        pass
    return render_to_response('login.html')


def create_wordsmith(request):
    if request.method == 'POST':
        pass
    return render_to_response('wordsmith_form.html')


def my_tomes(request):
    context = RequestContext(request)
    tome_controller = TomeController()
    wordsmith = tome_controller.get_wordsmith()
    context['tomes'] = tome_controller.get_tomes(wordsmith)

    return render_to_response('tomes.html', context=context)
