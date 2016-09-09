<div class="container">
    <div class="row">
		<section>
        <div class="wizard">
            <div class="wizard-inner">
                <div class="connecting-line"></div>
                <ul class="nav nav-tabs" role="tablist">

                    <li role="presentation" class="active">
                        <a href="#step1" data-toggle="tab" aria-controls="step1" role="tab" title="Step 1">
                            <span class="round-tab">
                                <i class="glyphicon glyphicon-wrench"></i>
                            </span>
                        </a>
                    </li>

                    <li role="presentation" >
                        <a href="#step2" data-toggle="tab" aria-controls="step2" role="tab" title="Step 2">
                            <span class="round-tab">
                                <i class="glyphicon glyphicon-scale"></i>
                            </span>
                        </a>
                    </li>
                    <li role="presentation" >
                        <a href="#step3" data-toggle="tab" aria-controls="step3" role="tab" title="Step 3">
                            <span class="round-tab">
                                <i class="glyphicon glyphicon-road"></i>
                            </span>
                        </a>
                    </li>

                    <li role="presentation" >
                        <a href="#complete" data-toggle="tab" aria-controls="complete" role="tab" title="Complete">
                            <span class="round-tab">
                                <i class="glyphicon glyphicon-ok"></i>
                            </span>
                        </a>
                    </li>
                </ul>
            </div>

                <form class="form-horizontal" method="POST">
                    {{ .xsrfdata }} 
                    <div class="tab-content">
                        <div class="tab-pane active" role="tabpanel" id="step1">
                            <div class="step1">
                                <div class="row">  
                                    <!-- TARGA --> 
                                    <div class="form-group col-md-4">                
                                        <label class="control-label" for="targa">Targa</label>                          
                                        <div >
                                        {{if .Errors.Targa}}  {{.Errors.Targa}}  {{end}}
                                        <input id="targa" name="targa" type="text" value="{{.AutomezzoDG.Targa}}" size="7" class="form-control" placeholder="Targa"/>
                                        </div> 
                                    </div> 
                                    <!-- DATA IN FLOTTA --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="dataInFlotta">Data entrata in flotta</label>                     
                                        <div>
                                        {{if .Errors.AutomezzoDG.DataInFlotta}}  {{.Errors.AutomezzoDG.DataInFlotta}}  {{end}}
                                        <input id="dataInFlotta" name="dataInFlotta" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" value="{{.AutomezzoDG.DataInFlotta}}" size="10" class="form-control input-md" />
                                        </div> 
                                    </div>

                                    <!-- DATA FINE FLOTTA -->            
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="dataFineFlotta">Data uscita da flotta</label>
                                        <div>
                                        {{if .Errors.AutomezzoDG.DataFineFlotta}}  {{.Errors.AutomezzoDG.DataFineFlotta}}  {{end}}
                                        <input id="dataFineFlotta" name="dataFineFlotta" type=date min=2000-01-01 datetime="2000-01-01 00:00:01" value="{{.AutomezzoDG.DataFineFlotta}}" size="10" class="form-control input-md" placeholder="GG/MM/AAAA"/>
                                        </div>
                                    </div>   
                                    <!-- TIPOVEICOLO --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="tipoteicolo">Tipo Veicolo:</label>
                                        <div>  
                                        <select id="tipoveicolo" name="tipoveicolo" type="number" >
                                                <option value="1" selected="selected">Autoveicolo</option>
                                                <option value="2">Camion</option>
                                                <option value="3">Ciclomotore Stato</option>
                                                <option value="4">Furgone</option>
                                                <option value="5">Pullman</option>
                                                <option value="6">Motoveicolo</option>
                                                <option value="7">Altro</option>
                                        </select>
                                        </div>
                                    </div>  
                                    <!-- CONDIZIONE --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="condizione">Condizione:</label>
                                        <div>  
                                        <select id="condizione" name="condizione" type="number" size="1">
                                            <option value="1" selected="selected">Buono stato</option>
                                            <option value="2">Cattivo stato</option>
                                            <option value="3">Discreto Stato</option>
                                            <option value="4">In attesa di alienzazione</option>
                                            <option value="5">In attesa di Assegnazione</option>
                                            <option value="6">In attesa di Riparazione</option>
                                            <option value="7">Non utilizzabile</option>
                                            <option value="8">Rubato</option>
                                            <option value="9">Alienato</option>
                                        </select>
                                        </div>
                                    </div>   
                                    <!-- CONDUCENTE --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="conducente">Conducente:</label>
                                        <div>    
                                        <select id="conducente" name="conducente" type="number" size="1">
                                            {{.Conducentelist}}
                                        </select>
                                        </div>
                                    </div>  
                                    <!-- SETTORE --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="settore">Settore:</label>
                                        <div>  
                                        <select id="usettore" name="settore" type="number" size="1">
                                            <option value="1">Food</option>
                                            <option value="2">Lavanderia</option>
                                            <option value="3">Pulizia</option>
                                            <option value="4">Marketing</option>
                                            <option value="5">Officina</option>
                                            <option value="6" selected="selected">Agenti</option>
                                            <option value="7">Direzione</option>
                                        </select>
                                        </div>
                                    </div>  
                                    <!-- IMPIEGO --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="impiego">Impiego:</label>
                                        <div>  
                                        <select id="impiego" name="impiego" type="number" size="1">
                                            <option value="1" selected="selected">Aziendale</option>
                                            <option value="2">Aziendale + Personale</option>
                                            <option value="3">Personale</option>
                                        </select>
                                        </div>
                                    </div>    
                                    <!-- NOTE --> 
                                    <div class="col-md-4 form-group">
                                        <label class="control-label" for="note">Note</label> 
                                        <div> 
                                        {{if .Errors.AutomezzoDG.Note}}  {{.Errors.AutomezzoDG.Note}}  {{end}}    
                                            <input id="note" name="note" type="text" value="{{.AutomezzoDG.Note}}" size="100" class="form-control input-md"/>
                                        </div>
                                    </div>   
                                                                                                                                                                                                        
                                </div>


                            </div>
                            
                        </div>
                        <div class="tab-pane" role="tabpanel" id="step2">
                            <div class="step2">                            
                                    <div class="row">
                                        
                                        <!-- NUMERO LIBRETTO -->               
                                        <div class="col-md-4 form-group">
                                            <label class="control-label" for="numeroLibretto">Numero libretto</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.NLibretto}}  {{.Errors.AutomezzoDT.NLibretto}}  {{end}}
                                                <input id="numeroLibretto" name="numeroLibretto" type="text" value="{{.AutomezzoDT.NLibretto}}" size="20" class="form-control input-md"/>
                                            </div>
                                        </div>      
                                        <!-- NUMERO TELAIO -->               
                                        <div class="col-md-4 form-group">
                                            <label class="control-label" for="numeroTelaio">Numero di telaio</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.NLibretto}}  {{.Errors.AutomezzoDT.NLibretto}}  {{end}}
                                                <input id="numeroTelaio" name="numeroTelaio" type="text" value="{{.AutomezzoDT.NLibretto}}" size="17" class="form-control input-md"/>
                                            </div>
                                        </div> 
                                        <!-- ANNO IMMATRICOLAZIONE -->               
                                        <div class="col-md-4 form-group">
                                            <label class="control-label" for="annoImmatricolazione">Anno immatricolazione</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.AnnoImmatricolazione}}  {{.Errors.AutomezzoDT.AnnoImmatricolazione}}  {{end}}
                                            <select name="annoImmatricolazione" id="annoImmatricolazione" class="birthdrop">
                                                    <option value="1980">1980</option>
                                                    <option value="1981">1981</option>
                                                    <option value="1982">1982</option>
                                                    <option value="1983">1983</option>
                                                    <option value="1984">1984</option>
                                                    <option value="1985">1985</option>
                                                    <option value="1986">1986</option>
                                                    <option value="1987">1987</option>
                                                    <option value="1988">1988</option>
                                                    <option value="1989">1989</option>
                                                    <option value="1990">1990</option>
                                                    <option value="1991">1991</option>
                                                    <option value="1992">1992</option>
                                                    <option value="1993">1993</option>
                                                    <option value="1994">1994</option>
                                                    <option value="1995">1995</option>
                                                    <option value="1996">1996</option>
                                                    <option value="1997">1997</option>
                                                    <option value="1998">1998</option>
                                                    <option value="1999">1999</option>
                                                    <option value="2000">2000</option>
                                                    <option value="2001">2001</option>
                                                    <option value="2002">2002</option>
                                                    <option value="2003">2003</option>
                                                    <option value="2004">2004</option>
                                                    <option value="2005">2005</option>
                                                    <option value="2006">2006</option>
                                                    <option value="2007">2007</option>
                                                    <option value="2008">2008</option>
                                                    <option value="2009">2009</option>
                                                    <option value="2010">2010</option>
                                                    <option value="2011">2011</option>
                                                    <option value="2012">2012</option>
                                                    <option value="2013">2013</option>
                                                    <option value="2014">2014</option>
                                                    <option value="2015">2015</option>
                                                    <option value="2016" selected="selected">2016</option>
                                                </select>
                                            </div>
                                        </div>
                                        <!-- MARCA -->               
                                        <div class="col-md-6 form-group">
                                            <label class="control-label" for="marca">Marca</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.Marca}}  {{.Errors.AutomezzoDT.Marca}}  {{end}}
                                                <input id="marca" name="marca" type="text" value="{{.AutomezzoDT.Marca}}" size="15" class="form-control input-md"/>
                                            </div>
                                        </div>  
                                        <!-- MODELLO -->               
                                        <div class="col-md-6 form-group">
                                            <label class="control-label" for="modello">Modello</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.Modello}}  {{.Errors.AutomezzoDT.Modello}}  {{end}}
                                                <input id="modello" name="modello" type="text" value="{{.AutomezzoDT.Modello}}" size="15" class="form-control input-md"/>
                                            </div>
                                        </div>                                                                           
                                        <!-- KW-->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="kw">Kw</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.Kw}}  {{.Errors.AutomezzoDT.Kw}}  {{end}}
                                                <input id="kw" name="kw" type="number" min="0" value="{{.AutomezzoDT.Kw}}" size="4" class="form-control input-md"/>
                                            </div>
                                        </div>                     
                                        <!-- CILINDRATA-->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="cilindrata">Cilindrata</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.Cilindrata}}  {{.Errors.AutomezzoDT.Cilindrata}}  {{end}}
                                                <input id="cilindrata" name="cilindrata" type="number" min="0" value="{{.AutomezzoDT.Cilindrata}}" size="6" class="form-control input-md"/>
                                            </div>
                                        </div>  
                                        <!-- CONSUMO TEORICO -->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="consumoTeorico">Consumo Teorico</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.ConsumoTeorico}}  {{.Errors.AutomezzoDT.ConsumoTeorico}}  {{end}}
                                                <input id="consumoTeorico" name="consumoTeorico" min="0" value="{{.AutomezzoDT.ConsumoTeorico}}" size="4" class="form-control input-md"/>
                                            </div>
                                        </div>                                      
                                        <!-- KM ANNUI -->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="kmAnno">Km Annui</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.KmAnno}}  {{.Errors.AutomezzoDT.KmAnno}}  {{end}}
                                                <input id="kmAnno" name="kmAnno" type="number" min="0" value="{{.AutomezzoDT.KmAnno}}" size="10" class="form-control input-md"/>
                                            </div>
                                        </div>  
                                        <!-- COSTO KM -->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="costoKm">Costo per km</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.CostoKm}}  {{.Errors.AutomezzoDT.CostoKm}}  {{end}}
                                                <input id="costoKm" name="costoKm" type="number" min="0" value="{{.AutomezzoDT.CostoKm}}" size="15" class="form-control input-md" pattern="^\\$?(([1-9](\\d*|\\d{0,2}(,\\d{3})*))|0)(\\.\\d{1,2})?$"/>
                                            </div>
                                        </div>  
                                        <!-- PNEUMATICI -->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="pneumatici">Pneumatici</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.Pneumatici}}  {{.Errors.AutomezzoDT.Pneumatici}}  {{end}}
                                                <input id="pneumatici" name="pneumatici" type="text "value="{{.AutomezzoDT.Pneumatici}}" size="15" class="form-control input-md"/>
                                            </div>
                                        </div>                                            
                                        <!-- CARBURANTE-->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="cilindrata">Carburante</label>  
                                            <div> 
                                            {{if .Errors.AutomezzoDT.Carburante}}  {{.Errors.AutomezzoDT.Carburante}}  {{end}}
                                            <select id="carburante" name="carburante" type="number" size="1">
                                                <option value="1" selected="selected">Benzina</option>
                                                <option value="2">Disel</option>
                                                <option value="3">Gas</option>                            
                                                <option value="4">Metano</option>
                                            </select>
                                            </div>
                                        </div>  
                                        <!-- NORMATIVA EURO -->               
                                        <div class="col-md-3 form-group">
                                            <label class="control-label" for="normativaEuro">Normativa Euro:</label>
                                            <div>  
                                            {{if .Errors.First}}  {{.Errors.First}}  {{end}}  
                                            <select id="normativaEuro" name="normativaEuro" size="1">
                                                    <option value="0">Euro 0</option>
                                                    <option value="1">Euro 1</option>
                                                    <option value="2">Euro 2</option>
                                                    <option value="3">Euro 3</option>
                                                    <option value="4">Euro 4</option>
                                                    <option value="5" selected="selected">Euro 5</option>
                                                    <option value="6">Euro 6</option>
                                            </select>
                                            </div>
                                        </div>                                        
                                        
                                        <div class="col-md-12 form-group accordion-content">
                                        <label for="exampleInputFile">Allegato</label>
                                            <input type="file" id="exampleInputFile">
                                        </div>  
                                                                        
                                    </div>
                                    
                                
                            </div>
                            
                        </div>
                        <div class="tab-pane" role="tabpanel" id="step3">
                            <!-- TIPO CONTRATTO -->               
                            <div class="form-group">
                                <label class="col-md-6 control-label" for="contratto">Tipo Contratto</label>  
                                <div> 
                                {{if .Errors.Contratto}}  {{.Errors.Contratto}}  {{end}}
                                <select class="col-md-6 control-label" id="contratto" name="contratto" size="1">
                                        <option value="0"selected="selected">Acquisto</option>
                                        <option value="1">Leasing</option>
                                        <option value="2">Noleggio</option>
                                </select>
                                </div>
                            </div>
                            <hr>

                            <div class="step33">
                                <!-- NUMERO CONTRATTO -->               
                                <div class="form-group">
                                    <label class="col-md-6 control-label" for="numeroContratto">Numero Contratto</label>  
                                    <div class="col-md-6"> 
                                    {{if .Errors.NContratto}}  {{.Errors.NContratto}}  {{end}}
                                        <input id="numeroContratto" name="numeroContratto" type="text" value="{{.NContratto}}" size="20" class="form-control input-md"/>
                                    </div>
                                </div>   
                                                  
        
                                

                        
                            </div>
                            
                        </div>
                        <div class="tab-pane" role="tabpanel" id="complete">
                            <div class="step44">
                                <h5>Whats Next</h5>
                                <hr>
                                <div id="accordion-container">
                                    <h2 class="accordion-header"> Your Assigned Branch Office & Consultant</h2>

                                    <div class="accordion-content">
                                    <label for="exampleInputFile">Allega Libretto</label>
                                        <input type="file" id="exampleInputFile">
                                    </div>
                                    <div class="accordion-content">
                                    <label for="exampleInputFile">Allega Contratto</label>
                                        <input type="file" id="exampleInputFile">
                                    </div>
                                </div>
                                <ul class="list-inline pull-right">
                                <li><button class="btn btn-primary pull-right" type="submit" class="pull-right" value="Inserisci" />Inserisci</button></li>
                            </ul>
                            </div>
                        </div>
                        <div class="clearfix"></div>
                    </div>
            </form>
        </div>
    </section>
   </div>
</div>
