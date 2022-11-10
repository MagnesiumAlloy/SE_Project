

function Stack() {
  this.items = [];
  // 添加一个新元素到栈顶位置。push()
  Stack.prototype.push = function (element) {
    this.items.push(element);
  }
  // 移除栈顶的元素pop()
  Stack.prototype.pop = function () {
    return this.items.pop();
  }
  //  返回栈顶的元素，不对栈做任何修改peek()
  Stack.prototype.peek = function () {
    return this.items[this.items.length - 1];
  }
  // 判断栈是否空isEmpty()
  Stack.prototype.isEmpty = function () {
    if (this.items.length == 0) {
      return true;
    } else {
      return false;
    }
  }
  // 返回栈里的元素个数size()
  Stack.prototype.size = function () {
    return this.items.length;
  }
  // 将栈结构的内容以字符形式返回toString()
  Stack.prototype.toString = function () {
    var str = '';
    for (var i = 0; i < this.items.length; i++) {
      str += this.items[i] + ' ';
    }
    return str;
  }
}

function readDir(isroot, ID, path, inBin) {
  $.ajax({
    url: "/readDir",
    type: "GET",
    dataType: "json",
    data: {
      UserId: dt * 1,
      path: path,
      ID: ID,
      isroot: isroot,
      inBin: inBin
    },
    success: function (res) {
      IsRoot = isroot
      Path = path
      InBin = inBin;
      document.getElementById('tmpPath').value = Path;
      showTable(res.data, isroot, res.ID, res.Path, inBin);
    }
  })
};


function ReadDirWithPop() {
  let isroot = IsRoot;
  let filepath = Path;
  let inBin = InBin;
  if (stackID.size() <= 1 || stackPath.size() <= 1) {
    return;
  }
  stackID.pop();
  stackPath.pop();
  readDir(isroot, stackID.peek(), stackPath.peek(), inBin);
};

function ReadDirWithPush(isroot, id, path, inBin) {
  stackID.push(id);
  stackPath.push(path);
  readDir(isroot, id, path, inBin);
};

function rTime(date) {
  var json_date = new Date(date).toJSON();
  return new Date(+new Date(json_date) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '')
}

function ConvertSize(x) {
  if (x < 1024) {
    return x.toFixed(2) + "B"
  } else {
    x /= 1024.0
    if (x < 1024) {
      return x.toFixed(2) + "KB"
    } else {
      x /= 1024.0
      if (x < 1024) {
        return x.toFixed(2) + "MB"
      } else {
        x /= 1024.0
        if (x < 1024) {
          return x.toFixed(2) + "GB"
        }
      }
    }
  }
  return ""
}
