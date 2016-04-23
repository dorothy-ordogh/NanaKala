function toSessionStorage(obj) {
	return JSON.stringify(obj);
}

function fromSessionStorage(str) {
	return JSON.parse(str);
}

function signOut() {
	sessionStorage.removeItem("user");
	location.reload();
}

function checkSessionStorage() {
	if (sessionStorage.user) {
		$("#signinout").replaceWith("<li id='signinout'><a href='signin.html' onclick='signOut()'>Sign Out</a></li>");
	}
}