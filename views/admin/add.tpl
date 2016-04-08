<div class="col-md-10 content">
        <div class="panel panel-default">
        <div class="panel-heading">
            <h1>Aggiungi Utente</h1>
        </div>
        <div class="panel-body">
                <div id="content">
                &nbsp;
                {{if .flash.error}}
                <h3>{{.flash.error}}</h3>
                &nbsp;
                {{end}}{{if .flash.notice}}
                <h3>{{.flash.notice}}</h3>
                &nbsp;
                {{end}}
                {{if .Errors}}
                {{range $rec := .Errors}}
                <h3>{{$rec}}</h3>
                {{end}}
                &nbsp;
                {{end}}
                <form method="POST">
                {{ .xsrfdata }}   
                    <table>
                    <tr><td>First: {{if .Errors.First}}{{.Errors.First}}{{end}}</td><td><input name="first" value="{{.User.First}}" size="30" /></td></tr>
                    <tr><td>Last:</td><td><input name="last" value="{{.User.Last}}" size="30" /></td></tr>
                    <tr><td>Email:  {{if .Errors.Email}}{{.Errors.Email}}{{end}}</td><td><input name="email" value="{{.User.Email}}" size="30" /></td></tr>
                    <tr><td>Password: {{if .Errors.Password}}{{.Errors.Password}}{{end}}</td><td><input name="password" value="{{.User.Password}}" size="30" /></td></tr>
                    <tr><td>&nbsp;</td></tr>
                    <tr>
                        <td>&nbsp;</td><td><input type="submit" value="Add" /></td>
                    </tr>
                    </table>
                        <a href="http://{{.domainname}}/admin/{{.parms}}" >Indietro</a>
                </form>
                </div>
        </div>
        </div>

</div>
