from django.shortcuts import render_to_response
from django.template import RequestContext
from django.urls import get_resolver
from MyJournal import urls

def home_page(request):
    context = RequestContext(request)
    resolver = get_resolver(urls)
    url_list = {}
    for view, regexes in resolver.reverse_dict.items():
        if isinstance(view, str):
            url_list[view] = regexes[1]
    context['urls'] = url_list
    print(context['urls'])
    return render_to_response('home_page.html', context)