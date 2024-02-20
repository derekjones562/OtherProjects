from django.http import HttpResponse
from rest_framework import generics
from rest_framework import permissions
from rest_framework import status
from rest_framework.decorators import api_view, permission_classes
from rest_framework.permissions import DjangoModelPermissionsOrAnonReadOnly, IsAuthenticatedOrReadOnly
from rest_framework.response import Response
from rest_framework.reverse import reverse

from apps.tome.models import Page, WordSmith, Tome
from apps.tome.permissions import IsWordSmithOrReadOnly, IsTomesWordSmithOrReadOnly
from apps.tome.serializers import PageSerializer, WordSmithSerializer, TomeSerializer

"""
@api_view(['GET', 'POST'])
@permission_classes((DjangoModelPermissionsOrAnonReadOnly, ))
def page_list(request, format=None):
    """
# List all code snippets, or create a new snippet.
"""
    if request.method == 'GET':
        pages = Page.objects.all()
        serializer = PageSerializer(pages, many=True)
        return Response(serializer.data)

    elif request.method == 'POST':
        serializer = PageSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


@api_view(['GET', 'PUT', 'DELETE'])
def page_detail(request, pk, format=None):
    """
# Retrieve, update or delete a code snippet.
"""
    try:
        page = Page.objects.get(pk=pk)
    except Page.DoesNotExist:
        return HttpResponse(status=status.HTTP_404_NOT_FOUND)

    if request.method == 'GET':
        serializer = PageSerializer(page)
        return Response(serializer.data)

    elif request.method == 'PUT':
        serializer = PageSerializer(page, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    elif request.method == 'DELETE':
        page.delete()
        return HttpResponse(status=status.HTTP_204_NO_CONTENT)
"""


@api_view(['GET'])
def api_root(request, format=None):
    return Response({
        'users': reverse('user-list', request=request, format=format),
        'snippets': reverse('snippet-list', request=request, format=format)
    })


class WordSmithList(generics.ListAPIView):
    queryset = WordSmith.objects.all()
    serializer_class = WordSmithSerializer


class WordSmithDetail(generics.RetrieveAPIView):
    queryset = WordSmith.objects.all()
    serializer_class = WordSmithSerializer


class TomeList(generics.ListCreateAPIView):
    permission_classes = (permissions.IsAuthenticatedOrReadOnly,)
    queryset = Tome.objects.all()
    serializer_class = TomeSerializer

  #  def perform_create(self, serializer):
 #       serializer.save()


class TomeDetail(generics.RetrieveUpdateDestroyAPIView):
    permission_classes = (permissions.IsAuthenticatedOrReadOnly,
                          IsWordSmithOrReadOnly)
    queryset = Tome.objects.all()
    serializer_class = TomeSerializer


class PageList(generics.ListCreateAPIView):
    permission_classes = (permissions.IsAuthenticatedOrReadOnly,)
    queryset = Page.objects.all()
    serializer_class = PageSerializer


class PageDetail(generics.RetrieveUpdateDestroyAPIView):
    permission_classes = (permissions.IsAuthenticatedOrReadOnly,
                          IsTomesWordSmithOrReadOnly)
    queryset = Page.objects.all()
    serializer_class = PageSerializer



