var Service = {
    load: function(cb) {
        ajaxCall("GET", "/service/todos/", cb);
    },
    create: function(item, cb) {
        ajaxCall("POST", "/service/todos/", cb, item);
    },
    update: function(item, cb) {
        ajaxCall("PUT", "/service/todos/" + item.id, cb, item);
    },
    delete: function(uuid, cb) {
        ajaxCall("DELETE", "/service/todos/" + uuid, cb);
    }
}
$.ajaxSetup({
    dataFilter: function(data, dataType) {
        if (dataType == 'json' && data == '') {
            return null;
        } else {
            return data;
        }
    }
});
var ajaxCall = function(method, path, cb, body) {
    $.ajax({
        type: method,
        url: path,
        data: body ? JSON.stringify(body) : null,
        success: cb,
        dataType: "json"
      });
}