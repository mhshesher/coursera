from django.http import HttpResponse

def index(request):
	return HttpResponse("Hi There! I am panda.\nThis is my 1st test application.")