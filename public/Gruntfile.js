module.exports = function (grunt) {
    require('time-grunt')(grunt);
    require('load-grunt-tasks')(grunt);

    var buildProperties = {
        appName: 'pubkeymanager',
        assetsPath: 'assets/',
        bowerFolder: '<%= config.assetsPath %>components/',
        jsFolder: '<%= config.assetsPath %>js/',
        cssFolder: '<%= config.assetsPath %>css',
        distFolder: '<%= config.assetsPath %>dist/',
        tplsFolder: '<%= config.assetsPath %>tpls/',
    };

    grunt.initConfig({
        pkg: grunt.file.readJSON('./package.json'),
        config: buildProperties,
        jshint: {
            options: {

                force: true,
                curly: true, // Require {} for every new block or scope.
                eqeqeq: false, // Require triple equals i.e. `===`.
                eqnull: true,
                latedef: false, // Prohibit variable use before definition.
                unused: false, // Warn unused variables.
                undef: true, // Require all non-global variables be declared before they are used.
                maxparams: 15,
                browser: true, // Standard browser globals e.g. `window`, `document`.
                validthis: true,
                globals: {
                    jQuery: true,
                    $: true,
                    angular: true,
                    alert: true,
                    console: true,
                    _: true,
                    NotificationFx: true,
                    Modernizr: true,
                    popup: true,
                    showNotAddedCartNotification: true,
                    showNotification: true,
                    self: true,
                    FB: true,
                    IosSlider: true,
                    sliderProductos: true,
                    backoffice: true,
                    safeAccess: true
                }
            },
            uses_defaults: ['!<%= config.jsFolder %>livereload.js', '<%= config.jsFolder %>**/*.js', '<%= config.jsFolder %>**/**/*.js']
        },
        watch: {
            css: {
                files: ['<%= config.cssFolder %>**/*.css'],
                tasks: ['default'],
                options: {
                    livereload: true
                }
            },
            js: {
                files: ['<%= config.jsFolder %>**/*.js'],
                tasks: ['default'],
                options: {
                    livereload: true
                }
            },
            html: {
                files: ['index.html', '<%= config.tplsFolder %>**/*.html'],
                tasks: ['default'],
                options: {
                    livereload: true
                }
            }
        },
        notify_hooks: {
            options: {
                enabled: true,
                max_jshint_notifications: 5, // maximum number of notifications from jshint output
                title: "<%= config.appName %> Project msg:"
            }
        },
        clean: {
            initclean: {
                src: ['<%= config.distFolder %>*.*']
            },
            postclean: {
                src: ['<%= config.distFolder %>*-libs.js']
            },
            postmin: {
                src: ['<%= config.distFolder %>*-libs.js', '<%= config.distFolder %>*.min.temp.js']
            }
        },
        html2js: {
            options: {
                base: 'assets/tpls',
                module: 'PubKeyManager.templates',
                singleModule: true,
                useStrict: true,
                htmlmin: {
                    collapseBooleanAttributes: true,
                    collapseWhitespace: true,
                    removeAttributeQuotes: true,
                    removeComments: true,
                    removeEmptyAttributes: true,
                    removeRedundantAttributes: true,
                    removeScriptTypeAttributes: true,
                    removeStyleLinkTypeAttributes: true
                },
                rename: function (moduleName) {
                    return '/' + moduleName.replace('.html', '');
                }
            },
            main: {
                src: ['<%= config.tplsFolder %>**/*.tpl.html'],
                dest: '<%= config.distFolder %>/templates.js'
            },
        },
        cssmin: {
            dist: {
                files: {
                    'assets/dist/style.min.css': [
                        'assets/components/bootstrap/dist/css/bootstrap.min.css',
                        'assets/components/font-awesome/css/font-awesome.min.css',
                        'assets/css/cosmo.bootstrap.min.css',
                        'assets/css/pubKeyManager.css',
                    ]
                }
            }
        },
        uglify: {
            options: {
                banner: '/*! <%= config.appName %> <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n',
                mangle: true,
                preserveComments: false,
                sourceMap: true,
                sourceMapIncludeSources: true
            },
            build: {
                files: [
                    {
                        src: [
                            '<%= config.bowerFolder %>angular/angular.min.js',
                            '<%= config.bowerFolder %>angular-route/angular-route.min.js',
                            '<%= config.bowerFolder %>angular-sanitize/angular-sanitize.min.js',
                            '<%= config.bowerFolder %>angular-bootstrap/ui-bootstrap-tpls.min.js',
                            '<%= config.bowerFolder %>angular-cookies/angular-cookies.min.js',

                            '<%= config.distFolder %>templates.js',


                            '<%= config.jsFolder %>pubkeymanager.js',
                            '<%= config.jsFolder %>/configs/routes.config.js',
                            '<%= config.jsFolder %>/constant/constants.js',
                            '<%= config.jsFolder %>/services/backend.service.js',
                            '<%= config.jsFolder %>/factories/httpInterceptor.factory.js',
                            '<%= config.jsFolder %>/controllers/navbar.ctrl.js',
                            '<%= config.jsFolder %>/controllers/login.ctrl.js',
                            '<%= config.jsFolder %>/controllers/register.ctrl.js',
                            '<%= config.jsFolder %>/controllers/keylist.ctrl.js',
                            '<%= config.jsFolder %>/controllers/key.modal.ctrl.js',
                            '<%= config.jsFolder %>/controllers/addkey.modal.ctrl.js',
                            '<%= config.jsFolder %>/controllers/delete.modal.ctrl.js',
                        ],
                        dest: '<%= config.distFolder %><%= config.appName %>.min.js'
                    }
                ]
            },
        },
    });

    grunt.loadNpmTasks('grunt-env');
    grunt.loadNpmTasks('grunt-contrib-jshint');
    grunt.loadNpmTasks('grunt-notify');
    grunt.loadNpmTasks('grunt-html2js');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-watch');

    // This is required if you use any options.
    grunt.task.run('notify_hooks');

    grunt.registerTask('default', ['clean:initclean', 'jshint:uses_defaults', 'html2js', 'cssmin:dist', 'uglify:build', 'clean:postmin']);

};