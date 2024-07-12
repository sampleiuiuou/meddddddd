<template>
    <div>
        <v-app-bar 
          app
          color="light-blue lighten-2"
          flat
          white_text
          
        >
          <v-avatar
            :color="$vuetify.breakpoint.smAndDown ? 'grey darken-1' : 'transparent'"
            size="32"
          ></v-avatar>
    
          <v-tabs
            centered
            class="ml-n9"
            color="white--text"
          >
            <v-tab 
            >
            Dashboard
            </v-tab>
            <v-tab 
            >
            stock-view
            </v-tab>
            <v-tab 
            >
            Bill entry
            </v-tab>
            
        </v-tabs>
          <v-btn>LOGOUT</v-btn>
        </v-app-bar>
    
           <v-container>
            <v-row>
              <v-col
                cols="12"
                sm="2"
              >
              </v-col>
    
              <v-col
                cols="12"
                sm="8"
              >
                <v-sheet
                  min-height="70vh"
                  rounded="lg"
                >
                  <!--  -->
                </v-sheet>
              </v-col>
    
              <v-col
                cols="12"
                sm="2"
              >
                <v-sheet
                  rounded="lg"
                  min-height="268"
                >
                  <!--  -->
                </v-sheet>
              </v-col>
            </v-row>
          </v-container>
        ></div>
    </template>
    
    <script>
      export default {
        data: () => ({
         
        }),
      }
    </script>
    
    
    