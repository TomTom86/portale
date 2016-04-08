<div id="content">
<h1>Remove User Account</h1>
&nbsp;
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{.flash.error}}.
            </div>
            {{end}}           
            {{if .Errors}}
            {{range $rec := .Errors}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{$rec}}.
            </div>      
            {{end}}
            &nbsp;
            {{end}} 

<div class="col-md-10 content">
        <div class="panel panel-default">
        <div class="panel-heading">
            Eliminazione Account
        </div>
        <div class="panel-body">
                <p><font size="3">Caution: all related transactions and data will also be removed. Are you sure?</font></p>
                <form method="POST">
                {{ .xsrfdata }}                
                <table>
                <tr>      
                    <td>Current password:</td>
                    <td><input name="current" type="password" /></td>
                </tr>
                <tr><td>&nbsp;</td></tr>
                <tr>
                    <td>&nbsp;</td><td><input type="submit" value="Remove" /></td><td><a href="http://localhost:8080/">Cancel</a></td>
                </tr>
                </table>
                </form>
        </div>
        </div>

</div>





