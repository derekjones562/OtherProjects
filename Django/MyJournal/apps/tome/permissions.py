from rest_framework import permissions

from apps.tome.models import WordSmith, Tome


class IsWordSmithOrReadOnly(permissions.BasePermission):
    """
    Custom permission to only allow WordSmiths of an object to edit it.
    """

    def has_object_permission(self, request, view, obj):
        # Read permissions are allowed to any request,
        # so we'll always allow GET, HEAD or OPTIONS requests.
        if request.method in permissions.SAFE_METHODS:
            return True

        # Write permissions are only allowed to the owner of the snippet.
        return obj.word_smith == WordSmith.objects.get(user=request.user)

class IsTomesWordSmithOrReadOnly(permissions.BasePermission):
    """
    Custom permission to only allow the WordSmith of a Tome to edit it.
    """

    def has_object_permission(self, request, view, obj):
        if request.method in permissions.SAFE_METHODS:
            return True
        word_smith = WordSmith.objects.get(user=request.user)
        return obj.tome_id == Tome.objects.get(word_smith=word_smith)