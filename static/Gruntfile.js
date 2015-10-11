module.exports = function(grunt) {
    grunt.initConfig({
	    pkg: grunt.file.readJSON('package.json'),

        concat: {
            options: {
                separator: ';',
            },            
            static_mappings: {
                files: [
                    {src: 'js/src/admin/*.js', dest: 'js/build/admin/<%= pkg.name %>.js'},
                    {src: 'js/src/site/*.js', dest: 'js/build/site/<%= pkg.name %>.js'},
                ],
            },
            
            dist: {
                src: ['js/src/admin/*.js'],
                dest: 'js/build/admin/<%= pkg.name %>.js',
            },
            dist: {
                src: ['js/src/site/*.js'],
                dest: 'js/build/site/<%= pkg.name %>.js',
            },
        },

	    uglify: {
	        options: {
		        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
	        },
            static_mappings: {
                files: [
                    {src: 'js/build/admin/<%= pkg.name %>.js', dest: 'js/build/admin/<%= pkg.name %>.min.js'},
                    {src: 'js/build/site/<%= pkg.name %>.js', dest: 'js/build/site/<%= pkg.name %>.min.js'},
                ],
            },
                
	        dist: {
		        src: ['js/src/admin/*.js'],
		        dest: 'js/build/admin/<%= pkg.name %>.min.js'
	        },
	        dist: {
		        src: ['js/src/site/*.js'],
		        dest: 'js/build/site/<%= pkg.name %>.min.js'
	        },

            dynamic_mappings: {
                files: [
                    {
                        expand: true, 
                        cwd: 'js/src/custom/',
                        src: ['**/*.js'], 
                        dest: 'js/build/custom/',
                        ext: '.min.js', 
                        extDot: 'first' 
                    },
                ],
            },
	    },
        
        sass: {
	        dist: {
		        files: {
		            'css/admin/admin.css'     : 'sass/admin/admin.scss',
		            'css/site/site.css'      : 'sass/site/site.scss',
		            'css/normalize.css' : 'sass/normalize.scss'
		        }
	        },
            dynamic_mappings: {
                files: [
                    {
                        expand: true, 
                        cwd: 'sass/custom/',
                        src: ['**/*.scss'], 
                        dest: 'css/custom/',
                        ext: '.css', 
                        extDot: 'first' 
                    },
                ],
            },
        },

	    watch: {
	        css: {
		        files: 'sass/**/*.scss',
		        tasks: ['sass']
	        },
	        js: {
		        files: 'js/**/*.js',
		        tasks: ['concat','uglify']
	        }
	    }

    });
    
    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-contrib-sass');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.registerTask('default',['concat','sass','uglify']);
}
