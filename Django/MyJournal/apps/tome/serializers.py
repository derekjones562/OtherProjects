from datetime import datetime

from django.contrib.auth.models import User
from rest_framework import serializers
from rest_framework.fields import DateField

from apps.tome.models import Page, WordSmith, Tome


class WordSmithSerializer(serializers.HyperlinkedModelSerializer):

    class Meta:
        model = WordSmith
        fields = ('url', 'user', 'first_name', 'last_name', 'tomes')


class TomeSerializer(serializers.HyperlinkedModelSerializer):
    pages = serializers.HyperlinkedRelatedField(many=True, view_name='page-detail', read_only=True)
    number_of_pages = serializers.IntegerField(read_only=True)

    class Meta:
        model = Tome
        fields = ('url', 'id', 'word_smith', 'tome_title', 'number_of_pages', 'pages')

    def create(self, validated_data):
        """
                Create and return a new `Tome` instance, given the validated data.
        """
        return Tome.objects.create(**validated_data)

    def update(self, instance, validated_data):
        """
               Update and return an existing `Tome` instance, given the validated data.
        """
        instance.word_smith = validated_data.get('word_smith', instance.word_smith)
        instance.tome_title = validated_data.get('tome_title', instance.tome_title)
        instance.number_of_pages = validated_data.get('number_of_pages', instance.number_of_pages)
        instance.save()
        return instance


class PageSerializer(serializers.HyperlinkedModelSerializer):

    class Meta:
        model = Page
        fields = ('url', 'id', 'page_id', 'page_date', 'tome_id', 'entry')

    def create(self, validated_data):
        """
                Create and return a new `Page` instance, given the validated data.
        """
        return Page.objects.create(**validated_data)

    def update(self, instance, validated_data):
        """
               Update and return an existing `Page` instance, given the validated data.
        """
        instance.page_id = validated_data.get('page_id', instance.page_id)
        instance.page_date = validated_data.get('page_date', instance.page_date)
        instance.tome_id = validated_data.get('tome_id', instance.tome_id)
        instance.entry = validated_data.get('entry', instance.entry)
        instance.save()
        return instance
