
  		<div class="col-md-10 content">
  			  <div class="panel panel-default">
                <div class="panel-heading">
                   <h1>View/Update</h1>
                </div>
                <div class="panel-body">
                                <div id="content">
                                <br>
                                {{if .flash.error}}
                                <h3>{{.flash.error}}</h3>
                                <br>
                                {{end}}
                                {{if .flash.notice}}
                                <h3>{{.flash.notice}}</h3>
                                <br>
                                {{else}}
                                <form method="POST">
                                {{ .xsrfdata }}                                
                                <table>
                                <tr><td>Id:</td><td><input name="id" value="{{.User.ID}}" size="10" readonly /></td></tr>
                                <tr><td>First: {{if .Errors.First}}{{.Errors.First}}{{end}}</td><td><input name="first" value="{{.User.First}}" size="30" /></td></tr>
                                <tr><td>Last:</td><td><input name="last" value="{{.User.Last}}" size="30" /></td></tr>
                                <tr><td>Email:  {{if .Errors.Email}}{{.Errors.Email}}{{end}}</td><td><input name="email" value="{{.User.Email}}" size="30" /></td></tr>
                                <tr><td>Password:</td><td><input name="password" value="{{.User.Password}}" size="30" readonly /></td></tr>
                                <tr><td>ID key</td><td><input name="idkey" value="{{.User.IDkey}}" size="50" /></td></tr>
                                <tr><td>Reg date</td><td><input name="regdate" value="{{.User.RegDate}}" size="50" readonly /></td></tr>
                                <tr><td>Reset key</td><td><input name="resetkey" value="{{.User.ResetKey}}" size="50" /></td></tr>
                                <tr><td>&nbsp;</td></tr>
                                <tr><td><input type="checkbox" id="delete" name="delete" />Delete</td></tr>
                                <tr><td>&nbsp;</td></tr>
                                <tr>
                                    <td>&nbsp;</td><td><input type="submit" value="Update" /></td>
                                </tr>
                                </table>
                                </form>
                                {{end}}
                                <h4><a href="http://{{.domainname}}/appadmin/index/{{.parms}}">Return to Index</a></h4>

                                <div id="dialog-confirm" style="display:none" title="Are you SURE?">
                                <p><span class="ui-icon ui-icon-alert" style="float:left; margin:0 7px 20px 0;"></span>This record will be deleted and all associated data will be lost.</p>
                                </div>

                                <script>
                                // the old way, using onclick=
                                function showDialog() {
                                    var r = confirm("Are you SURE?");
                                    if (r != true) {
                                        document.getElementById("delete").checked = false;
                                    }
                                }

                                // the new way, using jQuery
                                $(document).ready(function () {
                                    $("#delete").click(function () {
                                        if ($(this).is(":checked")) {
                                            $("#dialog-confirm").dialog({
                                                resizable: false,
                                                height:140,
                                                modal: true,
                                                buttons: {
                                                    "Okay": function() {
                                                        $(this).dialog("close");
                                                    },
                                                    Cancel: function() {
                                                        document.getElementById("delete").checked = false;
                                                        $(this).dialog("close");
                                                    }
                                                }
                                            });
                                        } else {
                                            $("#dialog-confirm").dialog("close");
                                        }
                                    });
                                });
                                </script>

                                </div>

                </div>
              </div>

  		</div>