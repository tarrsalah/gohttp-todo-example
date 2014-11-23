(function(ko,$) {
    function Task (data) {
	data = data || {};
	this.id = data.id || 0;
	this.todo = ko.observable(data.todo);
	this.done = ko.observable(data.done);
    };

    function TodoViewModel () {
	var that = this;

	that.task= new Task({
	    todo:"",
	    done: false
	});

	that.tasks= ko.observableArray();

	that.getTasks= function() {
	    $.getJSON("/api/tasks", function(data) {
		data.forEach(function(t) {
		    that.tasks.push(new Task(t));
		});
	    });
	};

	that.save= function() {
	    $.ajax({
		type: "POST",
		url: "/api/tasks",
		dataType: "json",
		data: ko.toJSON(that.task,"[todo, done]", 2)
	    }).done(function(data) {
		that.tasks.push(new Task(data));
	    }).fail(function(data) {
		that.getTasks();
	    });
	};

	that.update = function(task) {
	    $.ajax({
	    	type: "PUT",
	    	url: "/api/tasks/" + task.id,
	    	dataType: "json",
	    	data: ko.toJSON(task, "[id todo done]", 2)
	    }).fail(function(data)  {
		that.getTasks();
	    });
	    return true;
	};

	that.delete= function(task) {
	    $.ajax({
	    	type: "DELETE",
	    	url: "/api/tasks/" + task.id
	    }).done(function() {
		that.tasks.remove(task);
	    }).fail(function(data)  {
		that.getTasks();
	    });
	};

	that.getTasks();
    };

    ko.applyBindings(new TodoViewModel());
}(ko,jQuery));
