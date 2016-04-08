

<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h1>{{.UFirst}} {{.ULast}}</h1>
        </div>
        <div class="panel-body">
            <div id="content">

                {{if .flash.error}}
                &nbsp;
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
                {{end}}
                <form method="POST">
                    {{ .xsrfdata }}         
                    <!-- NOME --> 
                    <div class="form-group">                
                        <label class="col-md-10 control-label" for="first">Nome</label>  
                        <div class="col-md-10">
                            <input id="first" name="first" type="text" value="{{.UFirst}}" class="form-control input-md"/>
                        </div>
                    </div> 
                    <!-- COGNOME --> 
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="last">Cognome</label>
                        <div class="col-md-10">
                            <input id="last" name="last" type="text" value="{{.ULast}}" class="form-control input-md"/>
                        </div>
                    </div> 
                    <!-- EMAIL -->            
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="email">Email</label>
                        <div class="col-md-10">
                            <input id="email" name="email" type="text" value="{{.UEmail}}" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- LVL --> 
                    <div class="form-group">
                        <label class="col-md-5 control-label" for="userlvl">Privilegi Utente:</label>
                            <select id="userlvl" name="userlvl" size="1">
                                {{.Userlvllist}}
                            </select>
                    </div>  
                      <!-- BLOCCO --> 
                    <div class="form-group">
                        <label class="col-md-5 control-label" for="blocco">Utente Bloccato:</label> 
                            {{.Checkbloccato}}
                    </div>   
                    <!-- PASS -->                           
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="password">New password (must be at least 6 characters):</label>
                        <div class="col-md-10">
                            <input id="password" name="password" type="password" placeholder="Password" class="form-control input-md">
                            <span class="help-block">Inserisci una password di almeno 6 caratteri</span>
                        </div>
                    </div>  
                    <!-- PASS2 --> 
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="password2">Confirm new password:</label>  
                        <div class="col-md-10">                            
                            <input id="password2" name="password2" type="password" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- IDKEY --> 
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="idkey">ID key</label>  
                        <div class="col-md-10">
                            <input id="idkey" name="idkey" value="{{.IDkey}}" size="35" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- REGDATE --> 
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="regdate">Data di registrazione</label> 
                        <div class="col-md-10"> 
                            <input id="regdate" name="regdate" value="{{.RegDate}}" size="35" readonly class="form-control input-md"/>
                        </div>
                    </div>   
                    <!-- RESET KEY -->               
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="resetkey">Reset key</label>  
                        <div class="col-md-10">
                            <input id="resetkey" name="resetkey" value="{{.ResetKey}}" size="35" class="form-control input-md"/>
                        </div>
                    </div>  
                    <!-- STRUMENTI --> 
                    <div class="form-group">
                        <label class="col-md-10 control-label" for="apps">Strumenti:</label>  
                        <div class="col-md-5">
                            {{.Checkautomezzi}}
                        </div>                            
                        <div class="col-md-5">
                            {{.Checkservizi}}
                        </div>
                    </div>  
                    <!-- footer --> 
                    <div class="form-group ">
                        <div class="col-md-5 pull-left">
                        <a href="http://{{.domainname}}/admin/id!0!id__gte,0" >Indietro</a>
                        </div>
                        <div class="col-md-5 pull-right">
                        <input type="submit" value="Update" />
                        </div>          
                    </div>                              

                </form>
            </div>
        </div>
    </div>
</div>